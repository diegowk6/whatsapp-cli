// Package importer reads message history from the WhatsApp macOS app's SQLite
// databases and writes them into whatsapp-cli's message store.
package importer

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// appleEpochOffset is the difference between Apple's Core Data reference date
// (2001-01-01) and the Unix epoch (1970-01-01), in seconds.
const appleEpochOffset = 978307200

// DefaultMacOSSourcePath returns the default path to WhatsApp's macOS data directory.
func DefaultMacOSSourcePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return filepath.Join(home, "Library", "Group Containers",
		"group.net.whatsapp.WhatsApp.shared")
}

// MessageStore is the subset of store.MessageStore that the importer needs.
type MessageStore interface {
	StoreChat(jid, name string, lastMessageTime time.Time) error
	StoreMessage(id, chatJID, sender, content string, timestamp time.Time, isFromMe bool,
		mediaType, filename, url, directPath, mimeType string,
		mediaKey, fileSHA256, fileEncSHA256 []byte, fileLength uint64) error
}

// Progress is called periodically during import with running totals.
type Progress func(chats, messages int)

// ImportMacOS reads ChatStorage.sqlite from sourcePath and imports chats and
// messages into dst. It reports progress via the optional progress callback.
// Returns (chatsImported, messagesImported, error).
func ImportMacOS(sourcePath string, dst MessageStore, progress Progress) (int, int, error) {
	dbPath := filepath.Join(sourcePath, "ChatStorage.sqlite")
	if _, err := os.Stat(dbPath); err != nil {
		return 0, 0, fmt.Errorf("ChatStorage.sqlite not found at %s: %w", dbPath, err)
	}

	// Open read-only so we never corrupt the live WhatsApp database.
	dsn := fmt.Sprintf("file:%s?mode=ro&_foreign_keys=off", dbPath)
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return 0, 0, fmt.Errorf("open ChatStorage.sqlite: %w", err)
	}
	defer db.Close()

	chatsImported, err := importChats(db, dst, progress)
	if err != nil {
		return 0, 0, fmt.Errorf("import chats: %w", err)
	}

	messagesImported, err := importMessages(db, dst, progress)
	if err != nil {
		return chatsImported, 0, fmt.Errorf("import messages: %w", err)
	}

	return chatsImported, messagesImported, nil
}

func importChats(db *sql.DB, dst MessageStore, progress Progress) (int, error) {
	rows, err := db.Query(`
		SELECT
			ZCONTACTJID,
			COALESCE(ZPARTNERNAME, ZCONTACTJID),
			COALESCE(ZLASTMESSAGEDATE, 0)
		FROM ZWACHATSESSION
		WHERE ZCONTACTJID IS NOT NULL AND ZCONTACTJID != ''
	`)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var jid, name string
		var lastMsgApple float64
		if err := rows.Scan(&jid, &name, &lastMsgApple); err != nil {
			return count, err
		}
		lastMsg := appleToTime(lastMsgApple)
		if err := dst.StoreChat(jid, name, lastMsg); err != nil {
			return count, fmt.Errorf("store chat %s: %w", jid, err)
		}
		count++
		if progress != nil && count%100 == 0 {
			progress(count, 0)
		}
	}
	return count, rows.Err()
}

func importMessages(db *sql.DB, dst MessageStore, progress Progress) (int, error) {
	rows, err := db.Query(`
		SELECT
			COALESCE(m.ZSTANZAID, ''),
			c.ZCONTACTJID,
			COALESCE(m.ZFROMJID, ''),
			COALESCE(m.ZTOJID, ''),
			m.ZISFROMME,
			m.ZMESSAGETYPE,
			COALESCE(m.ZMESSAGEDATE, 0),
			COALESCE(m.ZTEXT, ''),
			COALESCE(mi.ZMEDIALOCALPATH, ''),
			COALESCE(mi.ZMEDIAURL, ''),
			COALESCE(CAST(mi.ZFILESIZE AS INTEGER), 0)
		FROM ZWAMESSAGE m
		JOIN ZWACHATSESSION c ON m.ZCHATSESSION = c.Z_PK
		LEFT JOIN ZWAMEDIAITEM mi ON m.ZMEDIAITEM = mi.Z_PK
		WHERE c.ZCONTACTJID IS NOT NULL
		ORDER BY m.ZMESSAGEDATE ASC
	`)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var stanzaID, chatJID, fromJID, toJID, text, mediaLocalPath, mediaURL string
		var isFromMe, msgType int
		var msgDateApple float64
		var fileSize int64

		if err := rows.Scan(
			&stanzaID, &chatJID, &fromJID, &toJID,
			&isFromMe, &msgType, &msgDateApple,
			&text, &mediaLocalPath, &mediaURL, &fileSize,
		); err != nil {
			return count, err
		}

		if stanzaID == "" {
			continue
		}

		ts := appleToTime(msgDateApple)
		sender := resolveSender(chatJID, fromJID, toJID, isFromMe == 1)
		mediaType, mimeType, filename := classifyMessage(msgType, mediaLocalPath)

		// For media messages without text, set a descriptive placeholder.
		content := text
		if content == "" && mediaType != "" {
			content = "[" + strings.Title(mediaType) + "]" //nolint:staticcheck
		}

		// Use the local path as both directPath (relative) and localPath hint.
		// We don't have encryption keys from the macOS app so downloads aren't
		// possible from whatsmeow, but the local file is often already on disk.
		localPathForStore := ""
		if mediaLocalPath != "" {
			localPathForStore = resolveMediaPath(mediaLocalPath)
		}

		if err := dst.StoreMessage(
			stanzaID, chatJID, sender, content, ts, isFromMe == 1,
			mediaType, filename, mediaURL,
			localPathForStore, // stored in direct_path column; no whatsmeow key available
			mimeType,
			nil, nil, nil, uint64(fileSize),
		); err != nil {
			return count, fmt.Errorf("store message %s: %w", stanzaID, err)
		}
		count++
		if progress != nil && count%1000 == 0 {
			progress(0, count)
		}
	}
	return count, rows.Err()
}

// maxValidAppleTimestamp is the Core Data timestamp for year 2100, used to
// clamp obviously corrupt values that appear in some WhatsApp records.
const maxValidAppleTimestamp = 3_155_760_000 // 2001-01-01 + 99 years

// appleToTime converts a Core Data timestamp (seconds since 2001-01-01) to time.Time.
// Timestamps outside [0, maxValidAppleTimestamp] are treated as zero (unknown).
func appleToTime(apple float64) time.Time {
	if apple <= 0 || apple > maxValidAppleTimestamp {
		return time.Time{}
	}
	unix := int64(apple) + appleEpochOffset
	return time.Unix(unix, 0).UTC()
}

// resolveSender returns the JID of the message sender.
func resolveSender(chatJID, fromJID, toJID string, isFromMe bool) string {
	if isFromMe {
		return "me"
	}
	if fromJID != "" && fromJID != chatJID {
		return fromJID
	}
	// For 1-on-1 chats ZFROMJID == ZCONTACTJID, so just use that.
	if fromJID != "" {
		return fromJID
	}
	return chatJID
}

// msgTypeToMedia maps WhatsApp macOS ZMESSAGETYPE integers to (mediaType, mimeType).
// Types observed in practice: 0=text, 1=image, 3=video/audio(opus), 6=voice,
// 7=link/extended-text, 8=document, 10=group-event, 14=contact-card,
// 15=gif/sticker. Others treated as text.
func classifyMessage(msgType int, localPath string) (mediaType, mimeType, filename string) {
	ext := strings.ToLower(filepath.Ext(localPath))

	switch msgType {
	case 1:
		mediaType = "image"
		switch ext {
		case ".png":
			mimeType = "image/png"
		case ".webp":
			mimeType = "image/webp"
		default:
			mimeType = "image/jpeg"
		}
	case 3:
		// In macOS app, type 3 covers both video and voice notes (opus).
		if ext == ".opus" || ext == ".ogg" || ext == ".m4a" || ext == ".mp3" {
			mediaType = "audio"
			mimeType = "audio/ogg; codecs=opus"
		} else {
			mediaType = "video"
			mimeType = "video/mp4"
		}
	case 6:
		mediaType = "audio"
		mimeType = "audio/ogg; codecs=opus"
	case 8:
		mediaType = "document"
		filename = filepath.Base(localPath)
	case 15:
		// Sticker/GIF — treat as image
		mediaType = "image"
		if ext == ".gif" {
			mimeType = "image/gif"
		} else {
			mimeType = "image/webp"
		}
	}
	return
}

// resolveMediaPath resolves a relative macOS media path to an absolute path
// inside the WhatsApp shared container, if it exists on disk.
func resolveMediaPath(relativePath string) string {
	if filepath.IsAbs(relativePath) {
		return relativePath
	}
	base := DefaultMacOSSourcePath()
	if base == "" {
		return relativePath
	}
	abs := filepath.Join(base, relativePath)
	if _, err := os.Stat(abs); err == nil {
		return abs
	}
	return relativePath
}
