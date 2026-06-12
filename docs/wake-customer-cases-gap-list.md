# Customer Cases — Gap List

**Discovery date**: 2026-06
**Method**: 3-phase systematic sweep — URL mining (`.com.br` + `.fbits.app`) + signal-phrase mining + cross-reference against existing 59 cases and 230 partners
**Source**: `store/messages.db` (165,240 group text messages)
**Existing catalog**: 59 cases (CASE-001 → CASE-059)

---

## Summary

| Bucket | Count |
|---|---|
| Phase 1A — distinct `.fbits.app` subdomains extracted | 19 (mostly partner sandboxes; ~3 customers) |
| Phase 1B — distinct `.com.br` domains extracted (after noise filter) | 491 |
| Phase 1B — domains in Wake-related groups (raw) | 282 |
| Phase 2 — signal-phrase hits (novo cliente, fechou, go-live, kickoff, onboarding, migrou, etc.) | 1,161 |
| Phase 2 — structured `Cliente:`/`Lojista:` field hits | 137 |
| Phase 3 — cross-reference rejections (partners, news, tools, internal, events) | ~190 |
| Phase 3 — duplicates of existing CASE-NNN | 9 |

**Net new candidates after triage**: ~130 worth tracking, of which ~42 high-confidence and ~33 medium-confidence have enough evidence for a full CASE-NNN write-up.

**Projected total customer catalog if all gaps written up**: 59 existing + ~75 high/medium = **~134 cases**.

---

## High-confidence new customers (3+ hits, clear context in Wake-partner groups)

| Customer | Hits | First | Last | Top groups | Vertical | URL | One-liner |
|---|---|---|---|---|---|---|---|
| Gregory | 8 | 2025-09-17 | 2026-05-18 | Tropa de Elite 3(2); Comercial Cross WAKE(2) | moda | `gregory.com.br` | Metakosmos provador virtual em produção |
| Avera / Averara | 7 | 2026-05-15 | 2026-06-01 | The Rise of Skywalker(3); WAKE   Time Parcerias(1) | moda | `averarabrand.com.br` | negociando prateleira infinita / POS |
| OffPremium | 5 | 2025-06-24 | 2025-10-29 | Wake + Marketplace Blueman(5) | varejo outlet | `offpremium.com.br` | Anymarket / Marketplace Blueman ativo |
| Eletro Mais (MT) | 5 | 2025-11-26 | 2025-12-10 | Duo e Wake Parcerias(3); Wake & A&EIGHT (B8One)(1) | varejo eletro | `eletromaismt.com.br` | lead via Duo — saindo de OpenCart |
| Cimafer | 5 | 2025-12-03 | 2025-12-10 | Duo e Wake Parcerias(3); Wake & A&EIGHT (B8One)(1) | construção/ferramentas | `cimafer.com.br` | lead via Duo, GMV R$1M/mês |
| Orit | 4 | 2025-09-17 | 2026-02-13 | Wake <> Ge Commerce(1); SAVVI & Wake(1) | moda intima | `orit.com.br` | lead em rotação multi-agência |
| Cerealista Express | 4 | 2025-11-27 | 2026-02-10 | Wake <> Ge Commerce(1); Wake <> Avocado Tech(1) | alimentos | `cerealistaexpress.com.br` | lead via Ge Commerce / Avocado Tech |
| Eletrônica Santana | 4 | 2025-09-18 | 2026-04-08 | Ecommerce House(3); Networki - E-commerce(1) | varejo eletro | `eletronicasantana.com.br` | lead via Ecommerce House |
| Inoxlon | 4 | 2026-02-12 | 2026-02-19 | Wake <> Engajar(2); Wake & A&EIGHT (B8One)(1) | industrial/aço | `inoxlon.com.br` | lead via Wake×Engajar e B8One |
| Singular Medicamentos | 4 | 2025-11-06 | 2026-04-24 | Wake <> Nairuz(1); Wake + N1(1) | farma | `singularmedicamentos.com.br` | lead via Nairuz / N1 / Cava Digital |
| Casafer | 4 | 2025-06-30 | 2025-06-30 | O2 Consultoria & Wake(2); Eficaz <> WAKE(2) | construção/varejo | `casafer.com.br` | lead via Eficaz, SAP, sai de WooCommerce |
| OG Brindes | 4 | 2025-06-24 | 2025-07-08 | Wake <> Série//A(1); Wake <> Econverse(1) | B2B brindes | `ogbrindes.com.br` | lead via FG / Eficaz / Econverse |
| SoulPink | 4 | 2025-08-26 | 2025-08-28 | SoulPink Geral(4) | beleza | `eusoulpink.com.br` | campanha ativa SoulPink Geral |
| Az Acessórios | 3 | 2025-06-23 | 2025-06-26 | Eficaz <> WAKE(2); Comercial Cross WAKE(1) | moda acessórios | `azacessorios.com.br` | case Eficaz |
| Dallanese Parafusos | 3 | 2025-06-25 | 2025-06-25 | Wake <> Ge Commerce(1); Wake & A&EIGHT (B8One)(1) | industrial/ferramentas | `dallanese.com.br` | lead multi-agência (Ge/B8One/Duo) |
| Balletto | 3 | 2025-07-11 | 2025-07-11 | Wake <> 3W(1); O2 Consultoria & Wake(1) | moda/calçados | `balletto.com.br` | proposta via O2 Consultoria |
| AR15 Sport | 3 | 2025-07-22 | 2025-09-12 | Wake <> Lemoon(1); Wake <> E-Plus(1) | esporte | `ar15sport.com.br` | lead via Lemoon / E-Plus / HERO |
| Campinas Diesel | 3 | 2025-08-12 | 2025-08-12 | Wake <> Catus(1); Eficaz <> WAKE(1) | auto/veículos | `campinasdiesel.com.br` | lead via Catus / Eficaz / Proex |
| 7 Oliveiras | 3 | 2025-08-18 | 2025-08-18 | Simples Inovação & Wake Parcerias(1); FZ COMMERCE & Wake(1) | ind/alimentos | `7oliveiras.com.br` | lead via FZ Commerce |
| Conceito Rouparia | 3 | 2025-09-03 | 2025-09-03 | Simples Inovação & Wake Parcerias(1); FZ COMMERCE & Wake(1) | moda | `conceitorouparia.com.br` | lead via FZ Commerce |
| Carinion Mix | 3 | 2025-09-15 | 2025-09-15 | Wake <> E-Plus(1); Parcerias Dotkom &  Wake(1) | ? | `carinonmix.com.br` | lead via E-Plus / Eficaz / Dotkom |
| Believe Professional | 3 | 2025-09-22 | 2025-09-22 | Simples Inovação & Wake Parcerias(1); FZ COMMERCE & Wake(1) | beleza | `believepro.com.br` | lead via FZ Commerce |
| Tazin Cosméticos | 3 | 2025-09-24 | 2025-09-24 | Wake <> Lemoon(1); Simples Inovação & Wake Parcerias(1) | beleza | `tazincosmeticos.com.br` | lead via Lemoon / Simples / FZ |
| Regina Salomão | 3 | 2025-10-02 | 2025-10-02 | Wake <> Avocado Tech(1); Simples Inovação & Wake Parcerias(1) | moda | `reginasalomao.com.br` | lead via Avocado Tech / Simples / FZ |
| Loja do Detalhamento | 3 | 2025-10-29 | 2025-10-29 | Simples Inovação & Wake Parcerias(1); FZ COMMERCE & Wake(1) | auto/detailing | `lojadodetalhamento.com.br` | lead via Simples / FZ / Proex |
| SGCB Brasil | 3 | 2025-10-29 | 2025-10-29 | Simples Inovação & Wake Parcerias(1); FZ COMMERCE & Wake(1) | auto/detailing | `sgcbbrasil.com.br` | lead via Simples / FZ / Proex |
| Alfredo Meyer | 3 | 2025-11-25 | 2025-11-25 | FZ COMMERCE & Wake(1); Agencia Proex & Wake(1) | moda | `alfredomeyer.com.br` | lead via FZ / Proex / Cava |
| Festa Color | 3 | 2025-12-01 | 2025-12-01 | Wake+Omni360+Affys(1); Wake <> 3W(1) | festas/decoração | `festacolor.com.br` | lead via Omni360/Affys / 3W / Eficaz |
| Ivan Joias | 3 | 2025-12-01 | 2025-12-01 | FZ COMMERCE & Wake(1); Agencia Proex & Wake(1) | joalheria | `ivanjoias.com.br` | lead via FZ Commerce / Proex / Cava |
| Sustent | 3 | 2025-12-16 | 2025-12-16 | Wake+Omni360+Affys(1); FZ COMMERCE & Wake(1) | ? | `sustent.com.br` | lead via Omni360/Affys / FZ / Cava |
| Shop2gether (referência) | 3 | 2026-01-30 | 2026-03-30 | Wake <> Nairuz(1); Wake <> Avocado Tech(1) | moda | `shop2gether.com.br` | referência de feature mix & match |
| Zinnia Leather | 3 | 2026-02-02 | 2026-02-02 | Wake <> Engajar(1); Wake & A&EIGHT (B8One)(1) | moda/acessórios | `zinnialeather.com.br` | lead via Engajar / B8One / Cava |
| Cittius | 3 | 2026-02-09 | 2026-02-09 | Wake+Omni360+Affys(1); Wake <> 3W(1) | ? | `cittius.com.br` | lead via Omni360/Affys / 3W / FZ |
| Gold Prime Joias | 3 | 2026-02-13 | 2026-02-13 | Wake <> Ge Commerce(1); SAVVI & Wake(1) | joalheria | `goldprimejoias.com.br` | lead via Ge Commerce / SAVVI / DRIVEN |
| HV Sul | 3 | 2026-04-14 | 2026-04-14 | Wake <> Catus(1); Duo e Wake Parcerias(1) | ? | `hvsul.com.br` | lead via Catus / Duo / Avocado Tech |
| Inigual | 3 | 2026-04-22 | 2026-04-22 | Parcerias Dotkom &  Wake(1); FZ COMMERCE & Wake(1) | ? | `inigual.com.br` | lead via Dotkom / FZ / Eficaz |
| Fuel (UseFuel) | 3 | 2026-04-27 | 2026-06-02 | Wake & Metakosmos(1); Serasa & Wake(1) | ? | `usefuel.com.br` | lead via Metakosmos / Serasa / Duo |
| Ecler Beauty | 3 | 2026-04-29 | 2026-04-29 | SAVVI & Wake(1); O2 Consultoria & Wake(1) | beleza | `eclerbeauty.com.br` | lead via SAVVI / O2 / DRIVEN |
| Pisomeco | 3 | 2026-04-29 | 2026-04-29 | SAVVI & Wake(1); Quick Digital & Wake Parcerias(1) | construção | `pisomeco.com.br` | lead via SAVVI / Quick Digital / DRIVEN |
| Nano4you Brasil | 3 | 2026-05-28 | 2026-05-28 | Wake <> Nairuz(1); Wake + N1(1) | ? | `nano4you.com.br` | lead via Nairuz / N1 / Dotkom |
| Unipeças MA | 3 | 2026-05-29 | 2026-05-29 | Visie & Wake(1); SAVVI & Wake(1) | auto | `unipecasma.com.br` | lead via Visie / SAVVI / DRIVEN |
| Brofitwear | 3 | 2025-11-11 | 2026-05-19 | Wake <> Catus(2); Quick Digital & Wake Parcerias(1) | moda fitness | `brofitwear.com.br` | PRIMEIRO Go-Live da Quick Digital! |

---

## Medium-confidence (2 hits, some context)

| Customer | Hits | First | Last | Top groups | Vertical | URL | One-liner |
|---|---|---|---|---|---|---|---|
| Miss Victoria | 2 | 2025-07-07 | 2025-10-24 | FG + Wake(1); Comercial Cross WAKE(1) | moda | `missvicttoria.com.br` | lead via FG / Comercial Cross |
| Curavino | 2 | 2025-08-12 | 2026-02-27 | Comercial Cross WAKE(1); Agencia Virtualiti   Wake(1) | vinhos | `curavino.com.br` | lead via Comercial Cross / Virtualiti |
| Casa Simões | 2 | 2025-08-13 | 2025-08-13 | Wake <> Catus(2) | varejo | `casasimoes.com.br` | lead via Catus |
| Uniland | 2 | 2025-08-15 | 2025-08-15 | Wake <> 3W(1); Simples Inovação & Wake Parcerias(1) | ? | `uniland.com.br` | lead via 3W / Simples |
| Anhanguera Ferramentas | 2 | 2025-09-09 | 2026-02-27 | SAVVI & Wake(1); Agencia Virtualiti   Wake(1) | ferramentas | `anhangueraferramentas.com.br` | lead via SAVVI / Virtualiti |
| Imperial Ferramentas | 2 | 2025-09-12 | 2025-09-22 | WAKE   Time Agência(1); HERO & Wake(1) | ferramentas | `imperialferramentas.com.br` | lead via WAKE Time Agência / HERO |
| Voomp Creators | 2 | 2025-09-24 | 2026-01-12 | Ecommerce House(2) | ? | `voompcreators.com.br` | lead via Ecommerce House |
| GoPure | 2 | 2025-09-29 | 2025-09-29 | Simples Inovação & Wake Parcerias(1); FZ COMMERCE & Wake(1) | beleza/saúde | `gopure.com.br` | lead via Simples / FZ |
| Vital Bio | 2 | 2025-10-02 | 2025-10-02 | Simples Inovação & Wake Parcerias(1); FZ COMMERCE & Wake(1) | saúde/orgânicos | `vitalbio.com.br` | lead via Simples / FZ |
| Biggie (abiggie) | 2 | 2025-10-09 | 2025-10-09 | Wake <> Ge Commerce(1); O2 Consultoria & Wake(1) | ? | `abiggie.com.br` | Shopify→Wake lead via Ge Commerce, ERP Tiny |
| Berzerk | 2 | 2025-10-16 | 2025-10-16 | Ecommerce House(2) | ? | `berzerk.com.br` | lead via Ecommerce House |
| Armazém CEC | 2 | 2025-11-04 | 2025-11-04 | O2 Consultoria & Wake(1); DRIVEN & Wake(1) | ? | `armazemcec.com.br` | lead via O2 / DRIVEN — Eletro Mais group |
| QualyAir | 2 | 2025-11-12 | 2025-11-12 | Wake <> Série//A(2) | casa/clima | `qualyair.com.br` | lead via Série//A |
| Menegatti Couro | 2 | 2025-11-24 | 2025-11-24 | Wake <> Nairuz(1); Wake <> 3W(1) | moda couro | `menegatticouro.com.br` | lead via Nairuz / 3W |
| Made in Guarda | 2 | 2025-11-25 | 2025-11-25 | Wake+Omni360+Affys(1); Agencia Virtualiti   Wake(1) | ? | `madeinguarda.com.br` | lead via Omni360/Affys / Virtualiti |
| Loja Brasil Vida | 2 | 2025-11-27 | 2026-01-26 | Wake <> Avocado Tech(2) | ? | `lojabrasilvida.com.br` | lead via Avocado Tech |
| Magazine Claudia Lima | 2 | 2025-11-27 | 2025-11-29 | Ecommerce House(2) | ? | `magazineclaudialima.com.br` | lead via Ecommerce House |
| Magnani Embalagens | 2 | 2025-12-15 | 2025-12-15 | Ecommerce House(2) | embalagens | `magnaniembalagens.com.br` | lead via Ecommerce House |
| Você Constrói | 2 | 2025-12-15 | 2026-02-18 | Eficaz <> WAKE(2) | construção/B2B | `voceconstroi.com.br` | relacionado a CASE-051 Esplane |
| Camys | 2 | 2025-12-16 | 2026-02-27 | Mastercard & Wake(1); Agencia Virtualiti   Wake(1) | ? | `camys.com.br` | lead via Mastercard / Virtualiti |
| Puriflora | 2 | 2026-01-15 | 2026-01-15 | Wake+Omni360+Affys(1); Wake <> 3W(1) | ? | `puriflora.com.br` | lead via Omni360/Affys / 3W |
| Eleve Suplementos | 2 | 2026-01-15 | 2026-01-15 | Wake+Omni360+Affys(1); Wake <> 3W(1) | suplementos | `elevesuplementos.com.br` | lead via Omni360/Affys / 3W |
| Termomax | 2 | 2026-01-28 | 2026-01-28 | SAVVI & Wake(1); O2 Consultoria & Wake(1) | industrial | `termomax.com.br` | lead via SAVVI / O2 |
| Tanto (?) | 2 | 2026-02-05 | 2026-02-05 | SAVVI & Wake(1); O2 Consultoria & Wake(1) | ? | `tanto.com.br` | lead via SAVVI / O2 (verify name) |
| Dominik / Dominik Metalcenter | 2 | 2026-02-13 | 2026-02-19 | Wake <> Série//A(2) | industrial | `dominik.com.br` | lead via Série//A |
| Brasil Seikyo (verify) | 2 | 2026-02-18 | 2026-02-18 | O2 Consultoria & Wake(1); Duo e Wake Parcerias(1) | publicação? | `cile.brasilseikyo.com.br` | lead via O2 / Duo |
| GDom (verify) | 2 | 2026-02-26 | 2026-02-26 | Wake <> Catus(1); Duo e Wake Parcerias(1) | ? | `gdom.com.br` | lead via Catus / Duo |
| Casa Pedro | 2 | 2026-05-04 | 2026-05-06 | Wake <> Nairuz(1); Parcerias Dotkom &  Wake(1) | ? | `casaspedro.com.br` | lead via Nairuz / Dotkom |
| Crise CIA | 2 | 2026-05-26 | 2026-06-02 | The Rise of Skywalker(1); Comercial Cross WAKE(1) | ? | `crisecia.com.br` | lead via Comercial Cross / Tropa de Elite |
| Fallco | 2 | 2025-08-21 | 2025-08-21 | FZ COMMERCE & Wake(2) | ? | `fallco.com.br` | lead via FZ Commerce |
| Tjcc | 2 | 2025-09-04 | 2025-09-16 | SoulPink Geral(2) | ? | `congresso.tjcc.com.br` | Pessoal, bom dia!!! Vou passar os códigos pro TJCC ,tá bom?  https://www.sympla.com.br/eve |
| Empório Bud | 2 | 2025-07-06 | 2025-09-12 | Wake <> Oásis(1); Konduto & Wake(1) | ? | `emporiobud.com.br` | lead via Konduto introduction |
| Loja Orio do Peixe | 2 | 2025-08-04 | 2025-09-19 | WAKE   Time Agência(1); FG + Wake(1) | ? | `lojaoriodopeixe.com.br` | lead via FG / WAKE Time Agência |

---

## Low-confidence (1 hit — short context, worth tracking but light evidence)

Names only, sorted by date of mention. Listed for completeness — most are leads in early funnel or single-message references.

- **Tiinside** (2025-07-04, `tiinside.com.br`) — https://tiinside.com.br/03/07/2025/idworks-erp-tem-crescimento-de-mais-de-200-no-primeiro-
- **Maiscorreios** (2025-07-05, `maiscorreios.com.br`) — https://www.maiscorreios.com.br/  👀
- **Braavo** (2025-08-12, `lp.braavo.com.br`) — Pessoal muito bom dia,  Não sou muito ativo por aqui por conta de uma outra comunidade que
- **Startupi** (2025-08-18, `startupi.com.br`) — https://startupi.com.br/omnik-recebe-aporte-de-r-5-milhoes/
- **Laurenti** (2025-08-28, `laurenti.com.br`) — https://www.laurenti.com.br/
- **Doctorshoes** (2025-09-11, `doctorshoes.com.br`) — https://www.doctorshoes.com.br/
- **Trademaster** (2025-09-17, `trademaster.com.br`) — https://trademaster.com.br/
- **Precifica** (2025-10-06, `precifica.com.br`) — https://precifica.com.br/
- **Deloittecomunicacao** (2025-10-07, `lp.deloittecomunicacao.com.br`) — Oii pessoal! Tudo bem?   Nesta sexta-feira, dia 10/10, vamos ter um happy hour no nosso es
- **Gocase** (2025-10-10, `gocase.com.br`) — https://www.gocase.com.br/capinha-para-celular/iphone/
- **Fenalaw** (2025-10-14, `fenalaw.com.br`) — https://www.fenalaw.com.br/
- **Forumcomprar** (2025-10-15, `forumcomprar.com.br`) — COMPRAR - Fórum & Expo https://www.forumcomprar.com.br/2026/#  01 de Setembro de 2026   8h
- **Fenim** (2025-10-15, `fenim.com.br`) — Fenim Fashion 2026 Data: 27 a 30 de janeiro de 2026 https://fenim.com.br/feira-gramado/
- **Uniecommerceweek** (2025-10-15, `uniecommerceweek.com.br`) — Pro-magno centro de eventos. 1, 2 e 3 de julho de 2026. https://www.uniecommerceweek.com.b
- **Skeps** (2025-10-27, `skeps.com.br`) — Gosto deles https://skeps.com.br/beeno/
- **Biosphere** (2025-11-10, `biosphere.com.br`) — https://biosphere.com.br/
- **Premiummachine** (2025-11-18, `premiummachine.com.br`) — RESUMO:   https://www.premiummachine.com.br/ 2.7K DE visualização / são cliente Wix  Eles 
- **Home2Homestay** (2025-11-22, `home2homestay.com.br`) — https://www.home2homestay.com.br
- **Ceex** (2025-11-24, `ceex.com.br`) — @230502388777182 vai ter esse evento aqui essa semana: https://ceex.com.br/. Você quer ir?
- **Cajubrasil** (2025-11-25, `cajubrasil.com.br`) — Passando pra contar que iniciamos agora os descontos de * até 70% OFF *  Corre!!! Cupom JU
- **Wolycasa** (2025-11-25, `wolycasa.com.br`) — https://www.wolycasa.com.br/
- **Madeiramadeira** (2025-11-27, `madeiramadeira.com.br`) — https://www.madeiramadeira.com.br/ Olá pessoal. Para quem quiser comprar móveis, além do m
- **Casatema** (2025-11-27, `casatema.com.br`) — Se for móveis infantis, a Casatema é o melhor destino com descontos exclusivos, apenas no 
- **Hugoboss** (2025-11-28, `hugoboss.com.br`) — A melhor época do ano chegou  Temos produtos com até 70% de desconto, tanto masculino quan
- **Probel** (2025-11-28, `probel.com.br`) — *O ECOMMERCE HOUSE é parceiro Probel, então essa é pra vocês, membros da maior comunidade 
- **Ckgsbknowledge** (2025-11-30, `ckgsbknowledge.com.br`) — Bom dia!  https://ckgsbknowledge.com.br/o-gigante-inovador-do-seculo-21/
- **Venddor** (2025-12-05, `venddor.com.br`) — https://www.venddor.com.br
- **Zydon** (2025-12-05, `zydon.com.br`) — tem a https://www.zydon.com.br/ tambem
- **Brasilcoms** (2025-12-08, `brasilcoms.com.br`) — Uma agência que criou esse ensaio  https://www.brasilcoms.com.br/
- **Bomme** (2025-12-19, `bomme.com.br`) — https://www.bomme.com.br
- **Moneri** (2025-12-21, `social.moneri.com.br`) — https://social.moneri.com.br/landing/PA51N36W  Esse é o link da Moneri para você se cadast
- **Gruposg3** (2026-01-02, `gruposg3.com.br`) — https://gruposg3.com.br/
- **Universalautomotive** (2026-01-05, `universalautomotive.com.br`) — Uma das maiores do Brasil: https://www.universalautomotive.com.br/
- **Meumulti** (2026-01-20, `meumulti.com.br`) — https://www.meumulti.com.br/#/aviso
- **Bethkalan** (2026-02-04, `bethkalan.com.br`) — Ajudando uma colega a divulgar algumas vagas  A vaga de Joias é analista de social media -
- **Monitormercantil** (2026-02-14, `monitormercantil.com.br`) — A justiça deferiu processamento do pedido de falência contra as principais empresas do gru
- **Constance** (2026-02-18, `constance.com.br`) — https://www.constance.com.br/
- **Pandorajoias** (2026-02-18, `pandorajoias.com.br`) — https://www.pandorajoias.com.br/
- **Maxtitanium** (2026-02-18, `secure.maxtitanium.com.br`) — https://secure.maxtitanium.com.br/
- **Searchlab** (2026-03-11, `searchlab.com.br`) — https://searchlab.com.br/p/bicicleta-eletrica-estudo-de-mercado-e-analise-de-buscas/
- **Supportelogistica** (2026-04-09, `supportelogistica.com.br`) — bom dia! alguém ja usou os servicos da Supporte Logística? https://supportelogistica.com.b
- **D4Sign** (2026-04-15, `secure.d4sign.com.br`) — Olá pessoal, tudo bem? @230502388777182 o pessoal do estadão enviou o contrato para formal
- **Printi** (2026-04-16, `printi.com.br`) — https://www.printi.com.br/
- **Nacar** (2026-04-19, `nacar.com.br`) — https://www.nacar.com.br/
- **F360** (2026-04-28, `f360.com.br`) — https://f360.com.br/
- **Curamed** (2026-04-28, `curamed.com.br`) — Fiz no lovable de teste e estudo https://curamed.com.br/
- **Realmovi** (2026-04-28, `realmovi.com.br`) — alguém já usou a https://www.realmovi.com.br/ para entregas?
- **Justapprove** (2026-04-30, `justapprove.com.br`) — https://justapprove.com.br/produtos/camiseta-huge-approve-x-gremio-ii-torcid-branco/
- **Shopee** (2026-05-05, `live.shopee.com.br`) — Olá! Junte-se a mim para PetCamp - Ofertas e Cupons na Shopee Live, em 5 de maio de 2026 1
- **Voxus** (2026-05-05, `voxus.com.br`) — Galera. Me pediram no off aqui depois dos materiais educacionais que mandei.  Fiz uma cura
- **Saopauloinnovationweek** (2026-05-05, `saopauloinnovationweek.com.br`) — Mari, o pessoal da Pulso Digital que fecha o estadão enviou esse convite aqui:   Oi! Tudo 
- **Washtoalhas** (2026-05-06, `washtoalhas.com.br`) — https://www.washtoalhas.com.br/jogo-de-toalha-banhao-4-pecas-guadalupe-taupe-e-azul   Boa 
- **Peticaopublica** (2026-05-06, `peticaopublica.com.br`) — https://peticaopublica.com.br/pview.aspx?pi=BR151434
- **Economicnewsbrasil** (2026-05-11, `economicnewsbrasil.com.br`) — Gente o mktplace Elo 7 fechou 😱https://economicnewsbrasil.com.br/2026/05/11/elo7-encerra-o
- **Mazurky** (2026-05-25, `mazurky.com.br`) — https://www.mazurky.com.br

---

## URLs in Wake-partner groups that map to ALREADY-CATALOGUED cases (duplicates / verify)

- `karcher-center-altex.com.br` (17 hits) — likely maps to an existing CASE-NNN. Verify and append URL to that case.
- `norterefrigeracao.com.br` (5 hits) — likely maps to an existing CASE-NNN. Verify and append URL to that case.
- `palaciodasessencias.com.br` (4 hits) — likely maps to an existing CASE-NNN. Verify and append URL to that case.
- `wallford.com.br` (3 hits) — likely maps to an existing CASE-NNN. Verify and append URL to that case.
- `loja.emporiodocelular.com.br` (3 hits) — likely maps to an existing CASE-NNN. Verify and append URL to that case.
- `plataformacvccorp.com.br` (2 hits) — likely maps to an existing CASE-NNN. Verify and append URL to that case.
- `lojaybera.com.br` (1 hits) — likely maps to an existing CASE-NNN. Verify and append URL to that case.
- `bawclothing.com.br` (1 hits) — likely maps to an existing CASE-NNN. Verify and append URL to that case.
- `yoraatacado.com.br` (1 hits) — likely maps to an existing CASE-NNN. Verify and append URL to that case.

---

## Signal-phrase Phase 2 extra candidates (from structured Cliente:/Lojista: fields not captured by URL mining)

These came from messages with `Cliente:` / `Lojista:` / `Loja:` / `Marca:` / `Razão Social:` field patterns. Many overlap with the URL list, but the following are unique mentions not yet covered:

- **Grupo Oásis** — ? — lead via Cliente: field
- **Babel-Azza** — ? — lead via Cliente: field
- **Cogem** — cooperativa — lead via Cliente: field — COOPERATIVA DE CREDITO COGEM
- **Adecil** — ? — lead via Cliente: field
- **Grupo Reluz** — ? — lead via Cliente: field
- **Empório das Toalhas** — casa — lead via Cliente: field
- **Madel Home Comércio de Móveis** — casa/móveis — lead via Cliente: field
- **Vou de Click Ltda** — ? — lead via Cliente: field
- **Livraria do IPP** — livraria — lead via Cliente: field
- **Atacadão Maromba** — B2B — lead via Cliente: field
- **Marabraz** — móveis/varejo — lead via Cliente: field
- **Dominik Metalcenter** — industrial — Cliente: field + URL
- **Home Sport Center** — esporte — lead via Cliente: field
- **Linha e Conforto** — casa/têxtil — lead via Cliente: field
- **Cacto Fitwear** — fitness — lead via Cliente: field
- **Anjorosa Consultoria Saúde** — saúde/consultoria — lead via Cliente: field
- **Central Acessório** — ? — lead via Cliente: field
- **Rabusch** — ? — lead via Cliente: field
- **MPS Distribuidora Mercantil** — B2B distribuição — lead via Cliente: field
- **GF10 Indústria e Comércio de Vestuário** — moda B2B — lead via Cliente: field
- **RENSZ Calçados** — calçados — lead via Cliente: field
- **Digitron** — ? — lead via Cliente: field
- **Secret Outlet** — varejo outlet — lead via Cliente: field
- **Rabittohs (pocmoda)** — moda — fbits.app subdomain — likely POC
- **Loja Slívia** — ? — fbits.app subdomain
- **Stayfilm Inc (referência)** — tool — Stayfilm is a Wake partner — likely partner-customer mix
- **People Interactive** — tool/agência — phrase mention
- **Magazord** — ? — phrase mention
- **Saurus** — ? — phrase mention
- **Duloren** — moda íntima — phrase mention
- **VNDA (referência)** — competitor platform — mention - 'devido eles' — likely a churn-source platform
- **Própria Day / Marca Própria Day** — evento — industry event — NOT customer
- **Snacks do Johnny** — alimentos snacks — phrase mention
- **Urban Helmets** — moto — phrase mention
- **Braavo** — ? — phrase mention
- **Benova** — ? — phrase mention — appears as agency partner

---

## Cross-reference rejections (candidate names rejected — not customers)

Names found by URL or signal-phrase mining that are NOT new customers. Documented to prevent re-discovery.

### Already in catalog (existing CASE-NNN)
- Karcher (CASE-003) — `karcher-center-altex.com.br`
- Empório do Celular (CASE-029) — `loja.emporiodocelular.com.br`
- Wallford (CASE-030) — `wallford.com.br`
- Norte Refrigeração (CASE-031) — `norterefrigeracao.com.br`
- BAW (CASE-017) — `bawclothing.com.br`
- Yamaha (CASE-037) — `yamahastore.com.br`
- Ybera (CASE-047) — `lojaybera.com.br`
- Palácio das Essências (CASE-048) — `palaciodasessencias.com.br`
- CVC Corp (CASE-052) — `plataformacvccorp.com.br`
- Yora (CASE-053) — `yoraatacado.com.br`
- Esplane / Você Constrói (CASE-051) — `voceconstroi.com.br`, `esplane.com.br`

### Wake partners (in catalog seed)
- Stone, Pagaleve, Mastercard, Konduto, Belluno, Stayfilm, Metakosmos, Anymarket, Leanwork, ClearSale, Mundo Verde — these are partners, not customers.
- Onfly, Asaas, BlueTicket, Mindsight, Equals, Parcelex, ControlID, Zoppy, Stocker — partner platforms/tools.
- Belluno Pag, Wave Commerce — partner orgs.
- Dotkom — agency partner.

### News/portal/event/conf sites
- `lupebar.com.br` — Wake event venue (HH bar)
- `cerbella.com.br` — poetry blog content shared in WSR group
- `essentialnutrition.com.br` — coupon shared in ECBR Wine Summit (not customer)
- `taqtaq.com.br`, `gaab.com.br`, `automyx.com.br`, `benway.com.br` — startup/event references
- `intermodal`, `fispalfoodservice`, `opoderdoentre`, `marcapropriaday`, `anamid`, `fiemg`, `seocamp` — events/awards/federations
- `glassdoor`, `eventbrite`, `lojaintegrada`, `plazasulshopping` — generic tools/sites
- `kabum.com.br` — referenced cupom, NOT a Wake customer
- `sanavita.com.br`, `shopclub.com.br` — referenced as employer / club, NOT customers
- `shiseidobrasil.com.br` — Wake internal bazar
- `osklen.com.br` — Metakosmos prospect mentioned alongside Wake, NOT a Wake customer (yet)
- `bobstore.com.br`, `richards.com.br`, `lafort.com.br`, `recco.com.br` — references in URL lists (current VTEX/Shopify customers being prospected by agencies)

---

## URLs found but no customer name yet identified

Single mentions in Wake-partner groups where the message was a bare URL with no context, or context was too truncated to extract a confident name. These deserve a manual follow-up sweep:

- `tudoentregue.com.br` (WAKE | Time Parcerias) — bare URL
- `ed3.com.br` (WAKE | Time Agência) — 'alguém conhece essa agencia?' → agency check, not customer
- `controlid.com.br`, `parcelex.com.br`, `equals.com.br`, `dguests.com.br` — partner suggestions, not customers
- `unabeleza.com.br` — Unikapro migrated to Nuvem → CHURN signal, worth a case write-up
- `wavecommerce.com.br` — agency partner referenced in WAKE Time Parcerias

---

## Notable findings worth a dedicated CASE-NNN write-up

Ranked by evidence strength and strategic relevance:

1. **Brofitwear** — *first Go-Live of Quick Digital agency*, 2026-05-19. Strong launch signal.
2. **OffPremium** — active Marketplace Blueman customer; Anymarket integration in flight.
3. **Gregory** — Wake + Metakosmos provador virtual case (production), moda.
4. **Avera / Averara** — active negotiation for prateleira infinita + POS integration question.
5. **Eletro Mais (MT) / Cimafer** — Duo Wake Parcerias lead duo, both saindo de OpenCart with R$1M GMV/mês claim.
6. **Casafer** — SAP+WooCommerce migration lead via Eficaz.
7. **Singular Medicamentos** — farma lead in multi-agency rotation (Nairuz, N1, Cava).
8. **OG Brindes** — B2B brindes, Diretor Geral contact captured.
9. **Biggie (abiggie)** — Shopify→Wake migration lead via Ge Commerce, ERP Tiny.
10. **Unikapro / Unabeleza** — *CHURN* signal (migrou para Nuvem). Worth documenting alongside Mpozenato/Ricardo Almeida/DLK pattern.
11. **Empório Bud** — Konduto-introduced lead with CNPJ + contact captured.
12. **Crise CIA** — recent (2026-06-02) Comercial Cross mention.
13. **Belluno + Wake** — has its own dedicated group, suggesting deeper integration story.
14. **Você Constrói group structure** — Esplane (CASE-051) is the B2B arm. Você Constrói itself + the broader group may merit a parent case.

---

## Recommended next pass

Given the volume:

- **42 high-confidence + 33 medium-confidence = 75 candidates** with enough evidence for a CASE-NNN write-up using the existing format.
- **55 low-confidence** candidates should be tracked but only written up if a second mention appears.
- **15 structured-field extras** (e.g., Atacadão Maromba, Cogem, Marabraz, Madel Home, MPS Distribuidora, RENSZ Calçados, GF10, Dominik Metalcenter, Empório das Toalhas, Vou de Click) likely correspond to leads worth ~5 line stubs each.

**Suggested workflow**:

1. Spawn a follow-up `wake-partnership-kb-writer` invocation with this gap list as input.
2. Batch write-up the 14 "notable findings" first (CASE-060 → CASE-073).
3. Add a second batch for the remaining high-confidence customers (CASE-074 → CASE-101 approx).
4. Track medium/low-confidence as 'watchlist' stubs that get promoted to full cases once they accumulate ≥3 message hits.

**Expected final catalog size**: ~120–135 cases (59 existing + ~75 new write-ups).

---

## Verificação

| Pass | Query | Result count |
|---|---|---|
| Phase 1A | `LIKE '%.fbits.app%'` (group msgs, no media) | 170 raw messages → 19 distinct subdomains |
| Phase 1B | `LIKE '%https://%' AND LIKE '%.com.br%'` | 123,840 raw rows → 643 distinct domains → 491 after noise filter → 282 in Wake-partner groups |
| Phase 2 | 25 signal phrases (novo cliente, fechou, kickoff, …) | 1,161 hits |
| Phase 2 (structured) | `LIKE '%Cliente:%' OR LIKE '%Lojista:%' …` | 137 hits → 64 distinct labelled names → 15 unique extras |
| Phase 3 (cross-ref) | Match against 59 cases + 230 partners + 16 internal/Wake teams | 9 duplicates, ~190 rejections, ~130 net new |

Filter applied throughout: `chat_jid LIKE '%@g.us'` and `(media_type IS NULL OR media_type='')`.
