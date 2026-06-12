# KB Extract — Projeto Serasa Pass <> Leanwork <> Wake
> JID: 120363422164488957@g.us | Extracted: 2026-06-04

---

## FAQ

---

### FAQ-1 — Qual é o campo `type` na criação de sessão e quando usar cada valor?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2025-09-01 |
| **Categoria** | `faq` |
| **Tags** | `api`, `session`, `type`, `desktop`, `mobile` |

**Pergunta:** A API de criação de sessão tem um campo `type`. Quais são os tipos disponíveis e quando usar cada um?

**Resposta:** O campo `type` determina o comportamento da sessão:
- **Type 1 (webhook / Desktop):** Usado quando o usuário está em um dispositivo desktop. O fluxo exibe um QR Code para que o usuário leia com o celular e faça o compartilhamento no app. Ao finalizar, a API bate no endpoint webhook configurado com o `sessionId`.
- **Type 2 (url / Mobile):** Usado quando o usuário está em um dispositivo móvel. O link da sessão redireciona diretamente para o app Serasa, sem mostrar QR Code. Ao finalizar, redireciona o usuário de volta à URL configurada.

O retorno da criação de sessão inclui um deeplink para o app, além da `webUrl` com o `sessionId`.

---

### FAQ-2 — Qual URL usar para o fluxo de biometria facial (desktop x mobile)?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2026-01-14 |
| **Categoria** | `faq` |
| **Tags** | `biometria`, `url`, `desktop`, `mobile`, `sesssao`, `deeplink` |

**Pergunta:** O link utilizado para biometria facial é `https://serasa.com.br/autorizacao/serasa-pass/verificacao/cpf?sessionId=###`. Esse link estava exibindo QR Code antes e agora manda direto para tela do CPF. O comportamento mudou?

**Resposta:** Sim, o comportamento mudou com a atualização de funcionalidades. A solução é **remover o `/cpf` da URL**:

```
INCORRETO: https://serasa.com.br/autorizacao/serasa-pass/verificacao/cpf?sessionId=###
CORRETO:   https://serasa.com.br/autorizacao/serasa-pass/verificacao?sessionId=###
```

Com a URL correta, no desktop aparece o QR Code para fazer o processo completo no celular antes do retorno do webhook. O `type` continua sendo enviado conforme o fluxo:
- Mobile: `type = url`
- Desktop: `type = webhook`

---

### FAQ-3 — Como funciona o formato dos webhooks para desktop e mobile?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2025-09-24 |
| **Categoria** | `faq` |
| **Tags** | `webhook`, `url`, `desktop`, `mobile`, `configuração` |

**Pergunta:** Como configurar as URLs de webhook e redirect para desktop e mobile? As URLs devem ser diferentes?

**Resposta:** Sim, devem ser configuradas URLs distintas para desktop e mobile:

- **Desktop (webhook de notificação):** `POST` para a URL com `sessionId` no body — a Serasa bate nesse endpoint quando o cliente terminar de compartilhar no app.
  - Exemplo: `https://loja.com.br/webhook/token/{sessionId}`

- **Mobile (URL de redirecionamento):** Após o compartilhamento no app, o usuário é redirecionado para essa URL — não é um POST, é um redirect.
  - Formato: `/webhook-mobile?sessionId={id}&eventName=user_consent`
  - A loja identifica o parâmetro e fecha a aba.

**Importante:** A URL de redirecionamento mobile deve ser configurada pelo lado da Serasa como destino após o compartilhamento. No caso do projeto modelo Wake, as URLs foram:
```
Desktop webhook: https://partnerserasa.fbits.store/webhook/token
Mobile redirect: https://partnerserasa.fbits.store/webhook-mobile/
```

---

### FAQ-4 — Como o QR Code funciona na loja? O usuário precisa do app Serasa instalado?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2025-09-03 |
| **Categoria** | `faq` |
| **Tags** | `qrcode`, `app`, `deeplink`, `mobile`, `câmera` |

**Pergunta:** Para ler o QR Code precisa de algum SDK ou funcionalidade especial no app? A câmera do celular padrão funciona?

**Resposta:** Sim, a câmera padrão do celular funciona para ler o QR Code, pois o código é um deeplink. Ao escanear o QR Code com qualquer câmera (ou leitor de QR padrão), o sistema operacional detecta que a URL é do app Serasa e redireciona automaticamente para o aplicativo, onde o usuário realiza o compartilhamento das credenciais. Não é necessário SDK adicional para a leitura do QR Code.

O app oficial da Serasa é o único necessário, disponível na App Store e Google Play.

---

### FAQ-5 — Quais eventos de monitoramento devem ser rastreados via JavaScript?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2025-09-22 |
| **Categoria** | `faq` |
| **Tags** | `monitoramento`, `analytics`, `javascript`, `eventos`, `tracking` |

**Pergunta:** Quais eventos de comportamento do usuário devem ser rastreados e enviados à API de monitoramento do Serasa?

**Resposta (confirmado pelo time Serasa):**

**Eventos essenciais:**
- Impressão (exibição do componente Serasa Pass)
- Clique no botão
- Visualização do QR Code
- Erro

**Metadados obrigatórios:**
- Tipo de dispositivo (desktop/mobile)
- Viewport
- User-Agent

**Eventos avançados (opcionais):**
- Tempo até o clique
- Porcentagem de visualização do botão

**Tecnologia recomendada:** `sendBeacon` (API nativa dos browsers), ideal para enviar dados de analytics sem bloquear a navegação.

A Serasa pode criar um endpoint específico para receber esses dados via `sendBeacon`.

---

### FAQ-6 — Quando deve ser disparado o endpoint de transação (API Transactions)?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2025-09-22 |
| **Categoria** | `faq` |
| **Tags** | `transactions`, `pagamento`, `checkout`, `api`, `processo` |

**Pergunta:** A API de Transactions deve ser chamada após a finalização do pedido (fechamento) ou após a aprovação do pagamento (confirmação)?

**Resposta:** O disparo deve ocorrer **após a aprovação do pagamento**, para que o meio de pagamento utilizado seja identificado. Isso garante visibilidade completa sobre a transação, incluindo o método de pagamento.

**Observação adicional:** Em pedidos com frete, foi necessário criar um endpoint adicional para contemplar o valor do frete no envio. O time de desenvolvimento (Igor) já havia implementado essa variação.

---

### FAQ-7 — Como funciona o acesso ao painel administrativo da loja modelo Wake (ambiente sandbox)?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2025-09-01 |
| **Categoria** | `faq` |
| **Tags** | `acesso`, `painel`, `ambiente`, `homologação`, `wake`, `credenciais` |

**Dados do ambiente de homologação:**

| Recurso | URL |
|---|---|
| Loja (Storefront) | `https://partnerserasa.fbits.store/` |
| Painel Admin | `https://partnerserasa.fbits.app/` |

**Formato de login:**
- Campo usuário: `partnerserasa\nome.sobrenome` (NÃO usar e-mail como login)
- Senha: Selecionar "Primeiro Acesso" para receber link de redefinição por e-mail

**Atenção:** O sistema bloqueia automaticamente usuários inativos após determinado prazo desde a criação ou reativação. Para reativar, contatar o administrador do painel.

---

### FAQ-8 — O campo `eventName` no redirect mobile retorna quais valores possíveis?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2026-02-02 |
| **Categoria** | `faq` |
| **Tags** | `eventName`, `webhook-mobile`, `querystring`, `frontend` |

**Pergunta:** No redirect mobile, o campo `eventName` na query string pode retornar valores diferentes de `user_consent`? Como tratar erros no frontend?

**Resposta:** O mapeamento dos possíveis valores do `eventName` está disponível no Swagger da API (enviado como arquivo ZIP por questões de segurança, não há página pública do Swagger). O campo pode retornar diferentes valores em caso de erro ou situações específicas, permitindo que o frontend trate cada caso adequadamente.

Formato do redirect mobile com `eventName`:
```
/webhook-mobile?sessionId={id}&eventName=user_consent
```

Em caso de sucesso, `eventName = user_consent` e `sessionId` contém o ID da sessão do usuário.

---

## Troubleshooting

---

### TROUBLESHOOTING-1 — Erro de login rápido no checkout da loja (configuração desabilitada)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2025-09-01 |
| **Categoria** | `troubleshooting` |
| **Tags** | `login-rápido`, `checkout`, `configuração`, `wake`, `bug` |

**Sintoma:** Erro no fluxo de login rápido no checkout da loja Wake.

**Causa:** Uma configuração estava desabilitada no backend da plataforma.

**Resolução:** A configuração foi reabilitada pelo time Wake. Após a correção, o fluxo de login rápido passou a funcionar corretamente.

**Aprendizado:** Sempre verificar configurações de plataforma quando o login rápido não funcionar — o problema pode não ser de código, mas de configuração de feature no painel da plataforma.

---

### TROUBLESHOOTING-2 — IPs bloqueados impedindo chamadas à API Serasa

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2025-09-02 |
| **Categoria** | `troubleshooting` |
| **Tags** | `ip`, `bloqueio`, `api`, `cors`, `middleware`, `firewall` |

**Sintoma:** Requisições do middleware da Leanwork para a API Serasa retornando erros. Nenhuma request chegando do lado Serasa.

**Causa:** Os IPs do servidor da Leanwork não estavam na whitelist do firewall/infra da Serasa.

**Resolução:**
1. IP específico `40.65.244.176` foi liberado como solução paliativa imediata
2. No dia seguinte, a liberação definitiva de todos os IPs foi realizada pelo time de infra cloud Serasa

**Processo de liberação:**
- Identificar os IPs de saída do ambiente do parceiro (Azure neste caso)
- Solicitar ao time de infra Serasa via canal interno
- Aguardar configuração — pode levar até 1 dia útil para liberação completa

---

### TROUBLESHOOTING-3 — Erro "Integration error" ao chamar `/v1/verifier/transaction/process`

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2025-10-28 |
| **Categoria** | `troubleshooting` |
| **Tags** | `transaction`, `process`, `api`, `erro`, `intermitência` |

**Sintoma:** Ao chamar o endpoint `/v1/verifier/transaction/process`, a API retorna:
```json
{
  "errors": "Integration error id: 964b26b1-8e71-4741-ba18-0f854f0060cd"
}
```

**Causa:** Instabilidade intermitente na infraestrutura da API Serasa. O fluxo retornava a funcionar sozinho após alguns minutos.

**Resolução:** Aguardar normalização. O time Serasa foi notificado para monitorar as intermitências e investigar a causa raiz.

**Observação:** Este mesmo padrão de intermitência se repetiu em outras ocasiões. Recomenda-se implementar retry com backoff exponencial ao chamar este endpoint.

---

### TROUBLESHOOTING-4 — Tela branca após compartilhamento no fluxo desktop (biometria)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2026-01-08 |
| **Categoria** | `troubleshooting` |
| **Tags** | `desktop`, `biometria`, `tela-branca`, `ux`, `webhook` |

**Sintoma:** No fluxo Desktop de biometria facial — após o usuário ler o QR Code com o celular, realizar a biometria e compartilhar os dados — a aba de compartilhamento que estava aberta no desktop ficava em branco.

**Causa:** Comportamento esperado era fechar ou notificar o usuário após o webhook ser recebido. A tela branca indicava que o webhook foi recebido mas o frontend não tratou o redirecionamento/fechamento adequadamente.

**Resolução:** Correção implementada pela Leanwork no frontend para tratar o retorno do webhook e redirecionar/fechar a aba corretamente. Os dados chegaram corretamente na Wake mesmo durante o período do bug.

---

### TROUBLESHOOTING-5 — Fluxo mobile não batia no webhook após compartilhamento

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2026-01-09 |
| **Categoria** | `troubleshooting` |
| **Tags** | `mobile`, `webhook`, `aba-suspensa`, `polling`, `redirect` |

**Sintoma:** No fluxo mobile (biometria via URL), após o compartilhamento no app Serasa, o webhook de notificação não era acionado. Os dados não chegavam na loja Wake.

**Causa raiz identificada:** Em mobile, a aba da Wake fica suspensa enquanto o usuário vai para o app Serasa. Mesmo que a Serasa bata no webhook, a aba suspensa não processa o evento.

**Soluções discutidas:**
1. **Polling no endpoint `/shared`:** A loja fica consultando periodicamente se as credenciais foram compartilhadas — elimina a dependência do webhook nesse fluxo.
2. **Redirect com queryString:** Ao finalizar o fluxo (sucesso ou falha), o app Serasa redireciona o usuário de volta à loja com informações via query string — também elimina o webhook.

**Resolução adotada:** Redirect com `sessionId` e `eventName` via query string na URL de retorno mobile. A URL de redirect foi atualizada para o formato:
```
/webhook-mobile?sessionId={id}&eventName=user_consent
```

---

### TROUBLESHOOTING-6 — Redirect mobile bloqueado pelo Safari como popup

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2025-10-14 |
| **Categoria** | `troubleshooting` |
| **Tags** | `safari`, `ios`, `popup`, `mobile`, `redirect`, `javascript` |

**Sintoma:** No iPhone com Safari, o redirect após compartilhamento mobile era bloqueado como se fosse um popup. O usuário não era redirecionado de volta à loja.

**Causa:** O componente de botão Serasa era implementado como `<button>` com `window.open()` via JavaScript, e o Safari por padrão bloqueia abertura de janelas/abas geradas por JavaScript.

**Resolução parcial:** No Chrome o comportamento foi corrigido. No Safari continuou com problemas.

**Workaround identificado:** Usar um link HTML nativo (`<a href="" target="_blank">`) em vez de JavaScript para a abertura. Links HTML nativos não são bloqueados pelo Safari como popup.

**Obstáculo:** O card do Serasa Pass era um `<button>` e as mudanças necessárias para desktop e mobile precisavam ser via JS. Refatoração foi necessária para resolver definitivamente.

**Lição aprendida:** Para links que precisam abrir em nova aba em ambientes móveis (especialmente iOS/Safari), sempre preferir `<a>` HTML nativo ao invés de `window.open()` via JavaScript.

---

### TROUBLESHOOTING-7 — E-mail de boas-vindas sendo disparado para clientes do evento

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2025-09-10 |
| **Categoria** | `troubleshooting` |
| **Tags** | `email`, `transacional`, `evento`, `loja-modelo`, `wake` |

**Sintoma:** Durante o evento Make Together (demonstração presencial), visitantes que cadastravam-se na loja modelo recebiam e-mail de boas-vindas da plataforma Wake — comportamento indesejado durante uma demo.

**Causa:** A loja modelo estava com os e-mails transacionais habilitados por padrão.

**Resolução:** O time Leanwork editou a referência do arquivo de template de e-mail no repositório Storefront. O ajuste foi publicado e verificado antes do evento. O e-mail de confirmação de compra continuou funcionando normalmente (comportamento desejado).

**Como fazer:** A desabilitação do e-mail de boas-vindas requer editar o arquivo HTML correspondente na pasta do Storefront no repositório Git — não é configuração do painel.

---

### TROUBLESHOOTING-8 — Erro na validação de e-mails no formulário de cadastro durante evento

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2025-09-10 |
| **Categoria** | `troubleshooting` |
| **Tags** | `email`, `validação`, `formulário`, `evento`, `domínio` |

**Sintoma:** Durante o evento, alguns e-mails de cadastro retornavam como inválidos mesmo sendo e-mails reais.

**Causa investigada:** O JavaScript de validação de e-mails na loja possivelmente rejeitava domínios desconhecidos ou de empresas (e.g., `usuario@empresa.com.br`). Domínios comuns (`@gmail.com`, `@hotmail.com`) funcionavam.

**Workaround durante o evento:** Usar e-mails de domínios conhecidos (Gmail, Hotmail) ao cadastrar durante a demonstração.

**Observação:** Pode existir uma política de bloqueio contra múltiplas tentativas de acesso do mesmo IP (proteção DDOS), que poderia ter contribuído com o comportamento.

---

## How-To

---

### HOW_TO-1 — Como configurar o middleware Wake-Serasa Pass em nova infraestrutura

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2026-04-30 |
| **Categoria** | `how_to` |
| **Tags** | `middleware`, `infra`, `deploy`, `azure`, `setup`, `onboarding` |

**Contexto:** Ao integrar Serasa Pass em nova loja Wake (para um cliente), o middleware precisa ser configurado em infraestrutura dedicada.

**Passos:**

1. **Obter o código-fonte:**
   - O repositório do middleware e frontend é gerenciado pela Leanwork
   - Para transferência para infra do cliente/Serasa, solicitar acesso via gestão da Leanwork
   - Compartilhamento por repositório Git (acesso por e-mail) ou pacote seguro

2. **Setup de infraestrutura (estimativa Leanwork):**
   - 5 dias para configuração de infra (Azure)
   - 2 dias para deploy das aplicações (back + front)
   - 1 dia para testes integrados

3. **Configurar URLs na Serasa:**
   - Webhook desktop: `https://{loja}/webhook/token` (POST com `sessionId` no body)
   - Redirect mobile: `https://{loja}/webhook-mobile?sessionId={id}&eventName=user_consent`
   - Informar as URLs ao time Serasa para cadastro no sistema

4. **Liberar IPs:**
   - Identificar IPs de saída da infraestrutura do parceiro
   - Solicitar liberação ao time de infra Serasa
   - Aguardar confirmação antes de iniciar testes de integração

5. **Configuração no painel Wake (por agência implementadora):**
   - Instalar o plugin Serasa Pass na loja
   - Para lojas Storefront: a agência tem autonomia para editar o template
   - Para lojas Style Editor: configuração diferente — validar com time Wake

6. **Teste final:**
   - Testar fluxo completo desktop (QR Code + compartilhamento no app)
   - Testar fluxo mobile (redirect direto para app)
   - Verificar recebimento dos webhooks
   - Validar preenchimento do formulário com dados retornados

---

### HOW_TO-2 — Como criar um usuário no painel administrativo Wake (ambiente parceiro)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2026-05-04 |
| **Categoria** | `how_to` |
| **Tags** | `acesso`, `painel`, `usuário`, `wake`, `admin` |

**Pré-requisito:** Ter acesso de administrador ao painel Wake.

**Passos:**
1. Acessar `https://partnerserasa.fbits.app/`
2. Qualquer usuário administrador pode criar novos usuários
3. No campo login, usar o formato: `partnerserasa\nome.sobrenome` (NÃO usar o e-mail)
4. Senha: o novo usuário deve clicar em "Primeiro Acesso" para receber link de redefinição por e-mail

**Formato de usuários:**
```
partnerserasa\nome.sobrenome
```

**Atenção:** Usuários inativos por muito tempo são bloqueados automaticamente. Para reativar, o administrador deve acessar o painel e reativar o usuário manualmente.

---

### HOW_TO-3 — Como implementar a detecção de dispositivo e criação de sessão com type correto

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2025-09-24 |
| **Categoria** | `how_to` |
| **Tags** | `session`, `type`, `device-detection`, `middleware`, `implementação` |

**Objetivo:** Criar a sessão Serasa Pass com o `type` correto baseado no dispositivo do usuário.

**Lógica de implementação:**

```javascript
// Detectar device
const isMobile = /Android|iPhone|iPad|iPod/i.test(navigator.userAgent);

// Criar sessão com type adequado
const sessionType = isMobile ? 2 : 1;

// Type 1 = Desktop (retorna QR Code + webhook ao final)
// Type 2 = Mobile (retorna redirect direto para app)
```

**Fluxo Desktop (type 1):**
- Criar sessão → receber `webUrl` com sessionId
- Exibir QR Code gerado a partir da URL
- Aguardar webhook POST em: `/{rota-configurada}` com `sessionId` no body
- Ao receber o webhook, buscar dados compartilhados via API de shared

**Fluxo Mobile (type 2):**
- Criar sessão → receber URL de redirect
- Redirecionar usuário para o app Serasa
- Aguardar redirect de volta para: `/webhook-mobile?sessionId={id}&eventName=user_consent`
- Ao receber o redirect, buscar dados compartilhados via API de shared

**Importante:** O fluxo deve ser executado sequencialmente. Não pular etapas — o app precisa validar a sessão antes de qualquer outra operação.

---

### HOW_TO-4 — Como testar o blocklist (bloqueio de CPF/dispositivo) no Serasa Pass

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2026-02-09 |
| **Categoria** | `how_to` |
| **Tags** | `blocklist`, `teste`, `cpf`, `dispositivo`, `segurança` |

**Objetivo:** Testar o comportamento da integração quando um CPF ou dispositivo está na blocklist do Serasa Pass.

**Passos:**
1. Fornecer ao time Serasa os dados do dispositivo/CPF a ser bloqueado:
   - CPF do usuário de teste
   - Modelo do celular
2. Aguardar confirmação do bloqueio pelo time Serasa
3. Executar o fluxo completo de compartilhamento com o CPF/dispositivo bloqueado
4. Verificar se o comportamento de erro é tratado adequadamente pela integração

**Tempo de execução do bloqueio:** Geralmente realizado no mesmo dia pelo time Serasa.

---

### HOW_TO-5 — Como solicitar a API Key para nova integração de cliente

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2026-06-01 |
| **Categoria** | `how_to` |
| **Tags** | `api-key`, `onboarding`, `cliente`, `ip`, `configuração` |

**Processo para obter API Key para um novo cliente:**

1. Enviar e-mail para o time Serasa com as seguintes informações:
   - IPs de saída da infraestrutura do cliente (para whitelist no firewall)
   - URLs de webhook (desktop e mobile redirect)
   - Domínio da loja

2. O time Serasa configura:
   - Whitelist de IPs
   - URLs de redirect/webhook
   - Gera a API Key para o cliente

3. Após liberação dos IPs, realizar teste de autenticação (chamada à API deve retornar 200, não 401 Unauthorized)

4. Com a API Key configurada, o ambiente está pronto para a agência implementadora configurar o plugin no painel Wake

**Para validar se a liberação foi feita:** Fazer uma requisição de teste via Postman contra qualquer endpoint da API Serasa e confirmar que não retorna `401 Unauthorized`.

---

## Best Practices

---

### BEST_PRACTICES-1 — Separação de branches para demonstrações e desenvolvimento ativo

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2025-09-17 |
| **Categoria** | `best_practices` |
| **Tags** | `git`, `branches`, `demo`, `apresentação`, `gestão` |

**Contexto:** O projeto tinha uma loja modelo funcional com dados mock/fake (fase 1) que a Serasa queria usar para apresentações a clientes, enquanto o desenvolvimento da fase 2 continuava.

**Prática recomendada:**
- Manter uma **branch de demonstração** (Fase 1) separada, com link de preview estável, para apresentações a clientes e prospects
- A **branch principal** (main) continua evoluindo com as novas funcionalidades da fase 2
- Avisar o time de desenvolvimento antes de apresentações, para evitar que deployments da branch principal afetem a URL de demo
- Ao apresentar para clientes, compartilhar o link da branch de demonstração, não o da branch de desenvolvimento

**Benefício:** Clientes e prospects sempre veem a versão estável e apresentável, sem os riscos de funcionalidades em desenvolvimento ou parcialmente implementadas.

---

### BEST_PRACTICES-2 — Gestão de ambientes de testes vs. produção em projetos piloto

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2025-09-26 |
| **Categoria** | `best_practices` |
| **Tags** | `ambiente`, `homologação`, `produção`, `api`, `testes` |

**Contexto:** Durante o projeto piloto, a integração foi desenvolvida diretamente em produção da API Serasa, pois havia dificuldades para liberar ambiente de homologação.

**Lições aprendidas:**
1. **Confirmar sempre o ambiente:** Antes de iniciar testes, verificar explicitamente se as chamadas estão indo para HML ou PRD. Perguntar diretamente à equipe responsável.
2. **Priorizar homologação:** Mesmo que demore mais, usar ambiente de HML evita impactos em dados reais e clientes.
3. **Dados de teste em PRD:** Se precisar usar PRD para testes (como foi o caso), garantir que os dados de teste sejam claramente identificáveis e não poluam relatórios de produção.
4. **Freeze de deploy:** Estabelecer períodos de freeze antes de eventos ou apresentações importantes, como o time Wake fez antes do evento Make Together.

---

### BEST_PRACTICES-3 — Tratamento de dados de usuário na integração Wake-Serasa Pass

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2026-03-12 |
| **Categoria** | `best_practices` |
| **Tags** | `dados-usuário`, `formulário`, `email`, `genero`, `cadastro` |

**Regras e decisões tomadas no projeto:**

1. **Conflito de e-mail:** O e-mail digitado pelo usuário no início do login simplificado NÃO deve ser substituído pelo e-mail retornado pela Serasa (credencial). A substituição causava erros de duplicidade de e-mail na plataforma.

2. **Campo Gênero:** Como a Serasa Pass não retorna informação de gênero, a decisão foi configurar o valor default como `"outros"` quando o campo for obrigatório na loja.

3. **Atualização de dados pós-compartilhamento:** Após o compartilhamento das credenciais, os dados do usuário na loja DEVEM ser atualizados com as informações retornadas pelo Serasa Pass. Isso é requisito do produto — sem a atualização, há problemas com parceiros que dependem dos dados validados.

4. **Desmembramento da credencial de contato:** A credencial de contato deve ser desmembrada (tratada separadamente) para evitar sobrescrever dados que o usuário já inseriu na loja.

---

### BEST_PRACTICES-4 — Arquitetura do middleware Wake-Serasa Pass

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2025-09-01 |
| **Categoria** | `best_practices` |
| **Tags** | `middleware`, `arquitetura`, `azure`, `backend`, `segurança` |

**Arquitetura utilizada no projeto piloto:**

O projeto implementou um **middleware** (backend intermediário) entre a loja Wake e a API Serasa, hospedado na Azure. O middleware tem as seguintes responsabilidades:

- **Autenticação:** Guarda as credenciais de acesso à API Serasa (API Key), nunca expostas no frontend
- **Gestão de sessões:** Cria sessões Serasa Pass e retorna a URL/QR Code ao frontend
- **Recebimento de webhooks:** Endpoint para receber notificações de compartilhamento da Serasa
- **Endpoint de user consent:** `/api/serasa/webhook-user-consent` para notificações do fluxo de compartilhamento

**Endpoint de webhook do middleware:**
```
POST https://webapp-middleware-wake-serasa-pass.azurewebsites.net/api/serasa/webhook-user-consent
Content-Type: application/json
{
  "sessionId": "1233231"
}
```

**Motivo do middleware:** A API Serasa requer whitelist de IPs. Centralizar as chamadas no middleware facilita a gestão de IPs e evita expor credenciais no frontend.

**Transferência de propriedade:** O middleware foi desenvolvido pela Leanwork e deve ser transferido para a infraestrutura do cliente/Serasa ao produtizar a integração. O código-fonte é entregue via repositório Git (acesso controlado por segurança).

---

### BEST_PRACTICES-5 — Modelo de sustentação e manutenção do plugin após entrega

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2026-03-23 |
| **Categoria** | `best_practices` |
| **Tags** | `sustentação`, `contrato`, `sla`, `plugin`, `pós-entrega` |

**Contexto:** Após a entrega do projeto piloto, a Serasa precisava de um modelo para continuar evoluindo e mantendo o plugin em produção para novos clientes Wake.

**Modelo adotado:**
- **Proposta de sustentação mensal:** Inclui horas/mês para melhorias contínuas, correção de bugs e suporte — obrigatório quando a Leanwork hospeda a infraestrutura
- **Setup único:** Cobrado uma única vez por ambiente configurado (não recorrente)
- **Infra:** Cobrada mensalmente (plano Básico ou Recomendado); valores em moeda estrangeira sujeitos a variação cambial

**Processo de aprovação interna Serasa:**
- Proposta passa pelo time de compras (PR)
- Aprovação financeira necessária para setup e mensalidade
- Contrato/aditivo formal antes de qualquer trabalho adicional

**Para novos clientes Wake:**
- A Serasa é responsável por indicar o cliente à Leanwork
- A agência implementadora do e-commerce (ex.: b8one) é responsável pela configuração do plugin no painel Wake
- A Leanwork cuida do setup de infra e testes de integração
- Prazo estimado total (após aprovação): 8 dias (Leanwork) + prazo da agência

---

### BEST_PRACTICES-6 — Padrão de comunicação e gestão de projeto em parceria tripartite

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Serasa Pass <> Leanwork <> Wake |
| **JID** | `120363422164488957@g.us` |
| **Data** | 2025-09-03 |
| **Categoria** | `best_practices` |
| **Tags** | `gestão`, `comunicação`, `checkpoint`, `parceria`, `tripartite` |

**Modelo de comunicação adotado no projeto:**

1. **Daily checkpoint:** Reunião diária (10h30) com representantes dos três lados (Serasa, Leanwork, Wake) para alinhamento de status e bloqueios
2. **Grupo WhatsApp:** Canal para comunicações assíncronas, impedimentos urgentes e compartilhamento de arquivos
3. **E-mail:** Formalização de atas, propostas comerciais e itens que precisam de rastreabilidade
4. **Frequência adaptativa:** Nos momentos de menor atividade (análise/estimativa), o checkpoint foi suspenso e substituído por updates assíncronos no grupo

**Papéis e responsabilidades:**
- **Serasa:** Define os requisitos do produto, valida entregas, fornece API Key e documentação
- **Leanwork:** Desenvolvimento do middleware e plugin Wake, gestão de projeto técnico
- **Wake:** Suporte à plataforma, configurações do painel, acesso aos repositórios do Storefront

**Lição aprendida:** Em projetos com três empresas distintas, é crítico que cada parte tenha clareza sobre quem é responsável por cada item de infraestrutura e configuração. Ambiguidades sobre "quem cuida de quê" causaram atrasos repetidos no projeto.
