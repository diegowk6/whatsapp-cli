# KB Delivery — Batch 02
**Artigos**: 55–104 do registry  
**Fontes**: kb-anymarket-wake.md (FAQ-005 ao fim), kb-comercial-cross.md (primeiros 19 artigos)  
**Total**: 50 artigos  
**Status**: prontos para importação  
**Data**: 2026-06

---

### FAQ-005 — O ANYMARKET aceita contas do tipo "vendedor" com CPF para integração com Magalu?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2026-02-27 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-solutions`, `audience-cs` |

**Pergunta:** O ANYMARKET aceita contas do tipo "vendedor" com CPF para integração com Magalu?

**Resposta:** O tema foi escalado para o time de CS ANYMARKET para confirmação. Verificar disponibilidade com o time de parcerias ANYMARKET antes de iniciar o onboarding de um seller com esse perfil.

---

### FAQ-006 — Como funciona o markup adicional por canal no ANYMARKET?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2026-02-13 |
| **Categoria** | `faq` |
| **Tags** | `contrato-comercial`, `audience-cs` |

**Pergunta:** Como configurar markup adicional por canal no ANYMARKET?

**Resposta:**

1. Acessar o módulo **Marketplace** dentro do painel ANYMARKET.
2. Configurar o "markup adicional" para cada canal separadamente.
3. Exemplo: definir `1,2` para Mercado Livre Clássico e `1,4` para Mercado Livre Premium para cobrir comissão + taxas de ~20%.
4. Todo novo anúncio criado para aquele canal herdará automaticamente esse markup.

**Para atualizar anúncios já criados com markup 1.0:**
- Baixar a planilha de anúncios no ANYMARKET.
- Alterar somente o campo "Markup Adicional" na planilha.
- Reimportar a planilha.

---

### FAQ-007 — Como remover anúncios duplicados ou de produtos Full no ANYMARKET?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2026-03-25 |
| **Categoria** | `faq` |
| **Tags** | `marketing-eventos`, `audience-cs` |

**Pergunta:** Como remover anúncios duplicados ou de produtos Full no ANYMARKET?

**Resposta:** Para remover anúncios do ANYMARKET (ex.: anúncios Full importados que não devem ser gerenciados pelo hub):

1. Acessar **Transmissão > Anúncios** no painel ANYMARKET.
2. Filtrar pelo tipo desejado (ex.: anúncios fullfillment).
3. Selecionar os anúncios a serem removidos e escolher uma das opções:
   - **"Finalizar anúncios do ANYMARKET"** → Remove do ANYMARKET **e** finaliza o anúncio no Mercado Livre.
   - **"Remover anúncios do ANYMARKET"** → Remove **apenas** do ANYMARKET (anúncio permanece ativo no canal).

> ⚠️ Atenção ao escolher a opção: finalizar no canal é irreversível.

---

### FAQ-008 — O ANYMARKET tem integração com Época Cosméticos?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2025-11-17 |
| **Categoria** | `faq` |
| **Tags** | `ecossistema-parceiros`, `audience-solutions`, `audience-sales` |

**Pergunta:** O ANYMARKET tem integração com Época Cosméticos?

**Resposta:** **Sim.** O ANYMARKET possui integração com a Época Cosméticos. Confirmar disponibilidade com o time de CS ANYMARKET antes de incluir no escopo de um seller.

---

### FAQ-009 — Como funciona a limitação de um único parceiro/canal por conta no ANYMARKET?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2025-11-11 |
| **Categoria** | `faq` |
| **Tags** | `ecossistema-parceiros`, `audience-solutions`, `audience-cs` |

**Pergunta:** Como funciona a limitação de um único parceiro/canal por conta no ANYMARKET?

**Resposta:** Atualmente, o ANYMARKET suporta apenas **um único parceiro** para envio de pedidos por conta/canal. Isso significa que, se um seller tem Mercado Livre e Netshoes integrados via ANYMARKET, os pedidos de ambos os canais serão roteados para o mesmo parceiro (ERP).

**Solução de contorno (workaround):**
- Renomear o parceiro para "ANYMARKET" e identificar o marketplace real usando as **informações adicionais do pedido** ou os **metadados**.

**Roadmap:**
- A melhoria para suporte a múltiplos parceiros por conta/canal estava prevista para iniciar em **Q1** (item de maior prioridade na lista de produto ANYMARKET no período).

---

### FAQ-010 — Como atualizar informações de um produto na Wake que não refletiram no ANYMARKET?

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2026-02-27 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-cs` |

**Pergunta:** Como atualizar informações de um produto na Wake que não refletiram no ANYMARKET?

**Resposta:** Quando uma alteração de produto na Wake (ex.: peso, dimensões) não é refletida automaticamente no ANYMARKET:

**Verificação inicial:**
- No módulo da Wake dentro do ANYMARKET, verificar se a configuração de atualização está definida como **"Atualizar todos os dados dos produtos"**.

**Opção 1 — Gatilho manual:**
- Fazer qualquer alteração no produto dentro da Wake (ex.: editar a descrição) para disparar o gatilho de sincronização.

**Opção 2 — Reimportação individual:**
- No ANYMARKET, acessar o módulo da Wake → aba **"Importação de produtos"**.
- Informar o SKU individualmente e solicitar reimportação.
- Isso **não duplica** o produto e geralmente atualiza os dados.

> ⚠️ **Fotos/imagens:** Não são atualizadas automaticamente. Para atualizar imagens, a opção é remover o produto do ANYMARKET e importar novamente, ou cadastrar via tela/planilha.

---

## Troubleshooting — Resolução de Incidentes

---

### TRB-001 — Erro de autenticação intermitente ao conectar token Wake no ANYMARKET

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2025-08-22 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-cs`, `audience-solutions` |

**Sintomas:** Clientes recebem erro ao tentar autenticar o token Wake no painel ANYMARKET durante sessões de onboarding. O erro ocorre de forma intermitente — em alguns casos funciona na segunda tentativa. Casos observados: Linha e Conforto, Você Constrói, Blueman, Daju, Suggar.

**Causa raiz:** Em vários casos, o time Wake não havia preenchido o formulário de solicitação de ambiente antes de tentar a autenticação. Sem o formulário, o processo interno de vinculação do ANYMARKET não é disparado.

**Resolução:**
1. Verificar se o formulário de solicitação de loja foi preenchido: https://forms.office.com/r/23NsmfzG2B
2. Após o preenchimento, aguardar liberação pelo time ANYMARKET (geralmente no mesmo dia útil).
3. Se o problema persistir após a liberação, tentar a autenticação novamente — em vários casos o erro se resolve sem ação adicional.
4. Se ainda assim não funcionar, abrir ticket no suporte ANYMARKET com evidências do erro.

> ℹ️ Em alguns casos (ex.: cliente Daju), o time ANYMARKET reportou: "Não fizemos exatamente nada, colocamos o token e agora foi". Isso sugere que pode haver uma fila de processamento interna.

---

### TRB-002 — Pedidos importando com canal errado (todos aparecem como Magalu)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2025-11-11 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-cs`, `audience-solutions` |

**Sintomas:** (BackWash) Todos os pedidos da Netshoes e Mercado Livre estavam sendo importados como se fossem pedidos do Magalu.

**Causa raiz:** Limitação atual do ANYMARKET — só é possível ter **um único parceiro** configurado por conta. Quando o seller tem múltiplos canais, todos os pedidos são roteados para o único parceiro configurado.

**Resolução:**
- Renomear o parceiro para "ANYMARKET" (nome genérico).
- Identificar o canal de origem de cada pedido usando as **informações adicionais do pedido** ou os **metadados** disponíveis na API ANYMARKET.
- A solução definitiva (múltiplos parceiros por canal) estava em roadmap para Q1.

---

### TRB-003 — Erro 422 ao inserir dados do parceiro/canal do pedido

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2025-09-24 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-cs` |

**Sintomas:** (Suggar) Erro ao processar pedido:
```
422 - {"resultadoOperacao":false,"codigo":1001,"mensagem":"Falha ao inserir os dados do parceiro/canal do pedido"}
```

**Causa raiz:** [fonte original sem informação] — a causa foi identificada e tratada internamente pelo time ANYMARKET, mas o detalhe não foi registrado no histórico da conversa (retorno foi documentado no Slack).

**Resolução:** Abrir ticket no suporte ANYMARKET informando o código do erro e o número do pedido afetado para investigação.

---

### TRB-004 — Produtos com erro de vinculação por categoria duplicada na Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2025-09-01 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-cs` |

**Sintomas:** (iLoft) Produtos vinculados a categorias na Wake apresentando erro por "falta de vínculo" no ANYMARKET. Exemplo: SKU `CXT067PR`.

**Causa raiz:** O seller tinha **duas árvores de categoria duplicadas** na Wake. O ANYMARKET importou a árvore "errada" (a duplicata), resultando em produtos sem vínculo de categoria válido.

**Resolução:**
1. O cliente deve **corrigir o produto** na Wake, atribuindo-o à categoria correta.
2. **Excluir a categoria duplicada** da Wake.
3. Aguardar a sincronização do ANYMARKET.

---

### TRB-005 — Instabilidade global por falha nos servidores AWS

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2025-10-20 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-cs` |

**Sintomas:** Instabilidade generalizada nos serviços ANYMARKET (lentidão, falhas de integração, indisponibilidade de ambiente).

**Causa raiz:** Intermitências nos servidores AWS (Amazon Web Services) nos EUA. Impacto global afetando marketplaces, plataformas, ERPs e outros participantes do ecossistema de e-commerce.

**Resolução:** Acompanhar status nas páginas oficiais e aguardar restabelecimento. Comunicar clientes afetados.
- AWS Health: https://health.aws.amazon.com/health/status
- ANYMARKET Status Page: https://status.anymarket.com.br/

> ℹ️ Esse tipo de falha não tem correlação com janelas de manutenção programadas do ANYMARKET.

---

### TRB-006 — Bloqueio de painéis por pendência financeira

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2026-01-23 |
| **Categoria** | `troubleshooting` |
| **Tags** | `suporte-escalação`, `audience-cs`, `audience-pam` |

**Sintomas:** Todos os painéis ANYMARKET bloqueados simultaneamente.

**Causa raiz:** Pendências financeiras relacionadas ao relatório de faturamento, gerando bloqueio temporário pelo time financeiro do ANYMARKET.

**Resolução:**
1. O time de Parcerias Wake deve regularizar o relatório financeiro junto ao ANYMARKET.
2. Após regularização, solicitar o religamento dos acessos ao time financeiro ANYMARKET.
3. O desbloqueio pode levar algumas horas após a regularização.

> ⚠️ Em caso de bloqueio, **todos** os painéis de todos os sellers são afetados simultaneamente. Comunicar o time de CS imediatamente.

---

### TRB-007 — Imagens não atualizadas no ANYMARKET após importação

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2026-02-12 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-cs` |

**Sintomas:** Fotos/imagens inseridas na Wake não são atualizadas automaticamente no ANYMARKET, mesmo com a configuração de atualização de produtos ativa.

**Causa raiz:** Confirmado pelo time ANYMARKET: imagens não são atualizadas automaticamente como os demais atributos de produto.

**Resolução:**
- **Opção 1:** Remover o produto do ANYMARKET e importar novamente.
- **Opção 2:** Cadastrar/atualizar as imagens diretamente via tela ou planilha no ANYMARKET.

---

### TRB-008 — Pedidos duplicados importados do Mercado Livre

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2025-12-01 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-cs` |

**Sintomas:** (BackWash) Pedidos duplicados do Mercado Livre sendo importados.

**Causa raiz:** Causa provável — a data de início da integração não foi configurada corretamente, fazendo o ANYMARKET buscar pedidos históricos anteriores ao go-live.

**Resolução:** Verificar e corrigir a data de início da integração nas configurações do canal dentro do ANYMARKET. Abrir ticket no suporte para ajuste manual, se necessário.

---

### TRB-009 — Status de pedidos não atualizado no canal Privalia

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2025-12-02 |
| **Categoria** | `troubleshooting` |
| **Tags** | `suporte-escalação`, `audience-cs` |

**Sintomas:** (Inbrands/Ellus) Pedidos constando como "expedido" na Wake e como "aguardando rastreio" na Privalia. Status não propagado entre os sistemas.

**Causa raiz:** Ponto de configuração na integração identificado pelo suporte ANYMARKET. O processo de remediação foi:
1. Suporte reprocessou os pedidos afetados manualmente (retornar status e avançar).
2. O canal da Privalia recebeu a atualização corretamente após o reprocessamento.

**Resolução:**
- Abrir ticket no suporte Wake informando os pedidos afetados.
- O suporte Wake aciona o suporte ANYMARKET para reprocessamento.
- Em casos de go-live próximo: priorizar resolução antes da virada para evitar backlog.

---

### TRB-010 — Erro de autenticação SHEIN: campo retornando "ANYMARKET"

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2026-04-01 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-cs` |

**Sintomas:** (Anjus) Erro ao autenticar o canal SHEIN — um dos campos das credenciais retorna o valor "ANYMARKET" em vez do valor correto, impedindo a autenticação.

**Causa raiz:** Credenciais inválidas fornecidas pelo marketplace SHEIN.

**Resolução:**
1. O seller deve acessar a tela de credenciais no portal da SHEIN e solicitar **novas credenciais**.
2. Com as novas credenciais, tentar a autenticação no painel ANYMARKET novamente.
3. Se persistir o erro com as novas credenciais, acionar o suporte ANYMARKET para validação conjunta.

---

### TRB-011 — Erro de faturamento via XML: ANYMARKET não encontra XML no e-mail

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2026-03-19 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-cs`, `audience-solutions` |

**Sintomas:** O ANYMARKET não encontra o XML da NF-e no e-mail, mesmo após o pedido ter sido faturado na Wake com dados fiscais.

**Causa raiz:** Causa mais comum — o XML foi encaminhado ao e-mail **depois** que o pedido já atualizou para "Faturado" no ANYMARKET com os dados fiscais. O ANYMARKET buscou os dados antes de o XML chegar.

**Resolução:** Fluxo correto de faturamento via e-mail:
1. **Primeiro:** Enviar o XML para o e-mail do ANYMARKET.
2. **Depois:** Atualizar o pedido na Wake com os dados fiscais.

Essa sequência garante que o ANYMARKET encontre o XML ao processar a atualização do status.

**Para faturamento via ERP (API):**
- Endpoint: `v2/orders/{id}/nfe`
- Verificar o campo `NeedInvoicedXml` no pedido (`v2/orders/{id}`) para saber se o pedido exige envio de XML.
- Documentação: https://developers.anymarket.com.br/api/v2/ky473o3qk52qb-orders-id-nfe

---

## How-To — Processos e Procedimentos

---

### HTO-001 — Como agendar a primeira reunião de onboarding/migração de um seller

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2025-07-10 |
| **Categoria** | `how_to` |
| **Tags** | `onboarding`, `audience-cs` |

**Processo padrão:**

1. **Preencher o formulário de solicitação de loja:** https://forms.office.com/r/23NsmfzG2B
   - Incluir o máximo de detalhes no campo "Observações".
   - Sugerir 3 datas/horários disponíveis do cliente.

2. **Aguardar a criação do painel** pelo time ANYMARKET (geralmente 1 dia útil após o forms).

3. **Verificar disponibilidade do analista ANYMARKET** usando os links de booking:
   - Henrique Maia: https://outlook.office.com/bookwithme/user/7a734b912e9d4e65ab41b9195fc72b96%40db1.com.br?anonymous&ismsaljsauthenabled
   - Lucas Rodrigues: https://outlook.office.com/bookwithme/user/0ace20555a9e4de9ac208c7d16c16a5d%40db1.com.br/meetingtype/9QPMQRf5_UyOvZdCylTahw2?anonymous&ismsaljsauthenabled

4. **Enviar o convite** para o analista ANYMARKET e para o cliente.

5. **Confirmar painel criado** antes da reunião: verificar com o time ANYMARKET se o ambiente do seller está disponível.

**Nomenclatura padrão de convites:**
```
[Tipo] Anymarket | [Nome do Seller] | [Marketplace] | [Etapa]
```
Exemplos:
- `Migração Anymarket | Home Sport Center | Magalu | Integração de catálogo Wake`
- `Novo cliente Anymarket | Blueman | Riachuelo | Transmissão de anúncios`
- `Migração Anymarket | PER Automação | Magalu | Go Live`

---

### HTO-002 — Como realizar a autenticação da nova API Magalu (V2)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2026-02-02 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-cs`, `audience-solutions` |

**Contexto:** A Magalu migrou sua API para a V2. Todos os sellers Wake precisaram autenticar a nova versão até o prazo definido (11/02/2026).

**Passo a passo:**
- Artigo oficial com instruções: https://suporte.anymarket.com.br/pt-BR/support/solutions/articles/19000171434

**Verificação:** O time ANYMARKET pode verificar quais contas ainda não realizaram a autenticação acessando o painel administrativo. Acionar o CS ANYMARKET para checagem prévia ao prazo.

---

### HTO-003 — Como configurar o envio de XML/NF-e para pedidos no ANYMARKET via ERP

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2026-02-13 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions`, `audience-cs` |

**Para faturamento via e-mail:**
- Artigo: https://suporteanymarket.freshdesk.com/a/solutions/articles/19000137966?lang=pt-BR&portalId=19000043126

**Para faturamento via ERP (API):**
1. Verificar no pedido o campo `NeedInvoicedXml` via `GET v2/orders/{id}` — indica se o pedido exige XML.
2. Enviar o XML via `POST v2/orders/{id}/nfe`.
3. Documentação: https://developers.anymarket.com.br/api/v2/ky473o3qk52qb-orders-id-nfe

**Fluxo correto (via e-mail):**
1. Enviar o arquivo XML para o e-mail configurado no ANYMARKET **antes** de atualizar o status do pedido na Wake.
2. Após confirmação de recebimento do XML, atualizar o pedido na Wake com os dados fiscais.

---

### HTO-004 — Como realizar a migração de seller do Mercado Livre (FOC → ANYMARKET)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2026-01-07 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-cs`, `audience-solutions` |

**Para sellers com múltiplas contas no Mercado Livre:**
- Qualquer erro durante a migração pode causar queda de conta — processo exige cuidado.
- Para esses casos, o ANYMARKET oferece **migração assistida** (com custo adicional; exceções podem ser negociadas para clientes específicos).
- O CS ANYMARKET responsável é envolvido nas tratativas junto com o CS Wake.

**Materiais de apoio disponíveis:**
- Portal de treinamento ANYMARKET: https://escolademarketplace.com.br/anystart/
- Escola de Marketplace: https://escolademarketplace.com.br

**Sellers que podem fazer sozinhos (self-service):**
- Sellers com uma única conta no Mercado Livre e sem histórico complexo de anúncios.
- CS Wake acompanha e o ANYMARKET apoia via suporte por ticket.

---

### HTO-005 — Como configurar gestão de crossdocking no ANYMARKET

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2025-09-26 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-solutions`, `audience-cs` |

**Contexto (Hiper Equipamentos):** Seller com modelo de crossdocking precisava configurar a gestão de estoque/expedição adequada no ANYMARKET para integração com Magalu.

**Solução definida:** Gestão de crossdocking feita via planilha (orientação do analista ANYMARKET). O processo foi repassado ao cliente durante sessão de onboarding, com agenda de go-live marcada na sequência.

Para configurações específicas de crossdocking, agendar call técnica com o time de onboarding ANYMARKET para orientação personalizada.

---

### HTO-006 — Como liberar acesso de perfil no painel ANYMARKET para usuário que não visualiza determinadas telas

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2026-03-30 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-cs` |

**Sintoma:** Usuário do seller não visualiza determinadas telas no painel ANYMARKET.

**Resolução:**
1. Acessar **Acesso > Perfis de acesso** no painel ANYMARKET.
2. Identificar qual perfil o usuário está associado.
3. Ajustar as permissões do perfil para incluir as telas necessárias.

> Requer usuário com privilégio de administrador no painel.

---

### HTO-007 — Como aplicar regra de preço "custo de frete grátis" corretamente no ANYMARKET

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2026-02-26 |
| **Categoria** | `how_to` |
| **Tags** | `contrato-comercial`, `audience-cs` |

**Contexto (Eletrolar):** Regra de preço de "custo de frete grátis" estava configurada mas não sendo aplicada nos anúncios já publicados.

**Causa:** A regra estava configurada mas o botão **"Aplicar para transmissões"** não havia sido clicado.

**Resolução:**
1. Acessar a configuração da regra de preço no ANYMARKET.
2. Clicar no botão **"Aplicar para transmissões"**.
3. Isso aplica a regra retroativamente a todos os anúncios já publicados com a tag de frete grátis, ajustando o preço com a inclusão do custo de frete.

---

## Best Practices — Boas Práticas

---

### BPR-001 — Padronização de nomenclatura de convites de onboarding

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2025-08-25 |
| **Categoria** | `best_practices` |
| **Tags** | `onboarding`, `audience-cs` |

**Recomendação:** Padronizar o nome dos convites de reunião de onboarding no formato `[Tipo] Anymarket | [Nome Seller] | [Marketplace] | [Etapa]`.

**Exemplos:**
- `Migração Anymarket | PER Automação | Magalu | Vínculo de anúncios`
- `Migração Anymarket | Home Sport Center | Magalu | Integração de catálogo Wake`
- `Migração Anymarket | Centro Elétrico | Magalu | Go Live`
- `Novo cliente Anymarket | Fourleaf | Dafiti | Vínculo de anúncios`
- `Novo cliente Anymarket | Blueman | Riachuelo | Transmissão de anúncios`

**Rationale:** Facilita o acompanhamento de múltiplos clientes em onboarding simultâneo. Permite que lideranças (CS e Onboarding) identifiquem rapidamente a criticidade e etapa de cada cliente ao receber o convite, facilita priorização quando há conflito de agendas e melhora a organização do time de onboarding ANYMARKET.

---

### BPR-002 — Usar booking dos analistas para evitar conflitos de agenda

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2025-08-20 |
| **Categoria** | `best_practices` |
| **Tags** | `integração`, `audience-cs` |

**Recomendação:** Ao marcar reuniões de onboarding com clientes, **sempre usar os links de booking dos analistas ANYMARKET** em vez de combinar horários manualmente. Atentar para a duração dos slots: os slots do ANYMARKET são de **30 minutos** e reuniões com clientes Wake geralmente precisam de **1 hora** — comunicar esse ponto ao agendar.

**Links de booking disponíveis:**
- Henrique Maia (Onboarding ANYMARKET): https://outlook.office.com/bookwithme/user/7a734b912e9d4e65ab41b9195fc72b96%40db1.com.br?anonymous&ismsaljsauthenabled
- Lucas Rodrigues (Onboarding ANYMARKET): https://outlook.office.com/bookwithme/user/0ace20555a9e4de9ac208c7d16c16a5d%40db1.com.br/meetingtype/9QPMQRf5_UyOvZdCylTahw2?anonymous&ismsaljsauthenabled

**Rationale:** Combinar horários manualmente já causou conflitos de agenda — dois sellers agendados no mesmo horário ou reuniões mais longas que os slots adjacentes. Usar o booking oficial evita esses problemas e respeita a capacidade do time ANYMARKET.

---

### BPR-003 — Gerenciamento proativo de alterações de preço no ANYMARKET

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2026-02-23 |
| **Categoria** | `best_practices` |
| **Tags** | `contrato-comercial`, `audience-cs` |

**Recomendação:** Gerenciar todos os preços pelo ANYMARKET, nunca diretamente no canal. Práticas:
1. **Sempre gerenciar preços pelo ANYMARKET**, nunca diretamente no canal.
2. Configurar regras de preço com markup adequado para cada canal antes do go-live.
3. Ao detectar anúncios com desconto irregular, pausá-los imediatamente via ANYMARKET até a correção do markup.
4. Após a correção, remover os anúncios problemáticos no ANYMARKET e recriar com o markup correto.
5. Alertar o cliente no onboarding: qualquer alteração de preço deve ser feita pelo ANYMARKET, não diretamente no marketplace.

**Rationale:** Sellers que gerenciam preços via planilha ou diretamente no canal (sem usar o hub) correm risco de criar inconsistências de preço que resultam em anúncios com descontos acima do permitido ou vendas abaixo do custo.

---

### BPR-004 — Fluxo de escalação de tickets críticos com risco de churn

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2026-01-28 |
| **Categoria** | `best_practices` |
| **Tags** | `suporte-escalação`, `audience-cs` |

**Recomendação:** Quando um ticket não recebe retorno em prazo adequado e há risco real de churn, seguir fluxo de escalação:
1. Reforçar o ticket com novas evidências e data de abertura original.
2. Acionar o **coordenador de suporte ANYMARKET** diretamente no grupo, informando:
   - Nome do cliente
   - Número do ticket
   - Data de abertura
   - Impacto relatado pelo cliente
3. Se o coordenador não responder em tempo hábil, escalar para **gestão ANYMARKET** via e-mail ou via time de Parcerias.
4. Registrar **todas** as interações no grupo para manter o histórico e visibilidade.

Indicadores de escalação urgente: ticket aberto há mais de 5 dias úteis sem resolução; cliente reportando queda de vendas ou impossibilidade de operar; risco iminente de cancelamento do contrato.

**Rationale:** Tickets parados sem visibilidade já levaram a clientes relatando perda de faturamento e vendas paralisadas. Escalação estruturada, com histórico documentado no grupo, garante priorização e protege o relacionamento com o cliente antes que o risco de churn se materialize.

---

### BPR-005 — Responsabilidades Wake vs. ANYMARKET nas migrações

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2026-02-11 |
| **Categoria** | `best_practices` |
| **Tags** | `integração`, `audience-cs`, `audience-pam` |

**Recomendação:** Seguir a divisão de responsabilidades acordada entre Wake e ANYMARKET nas migrações:

**Clientes que Wake toca e ANYMARKET acompanha (migração de legado Wake):**
- CS Wake conduz o onboarding.
- ANYMARKET participa como suporte técnico especializado.
- Exemplos: Adecil, Eletrolar, Gringamx.

**Clientes que ANYMARKET toca e Wake acompanha:**
- Time de onboarding ANYMARKET conduz.
- CS Wake monitora e é ponto de contato com o cliente.
- Exemplo: Livo.

**Migrações de lista original Magalu (pendentes):**
- Conduzidas pelo ANYMARKET com acompanhamento do CS Wake responsável.
- Lista de sellers e CS responsáveis deve ser mantida atualizada em planilha compartilhada.

**Após conclusão das migrações:**
- Demandas técnicas migram para o fluxo padrão de tickets no portal de suporte.
- O modelo de acompanhamento assistido (dailys, onboarding conjunto) é suspenso.

**Rationale:** [fonte original sem informação] — a divisão foi acordada operacionalmente entre as equipes; o histórico da conversa não registra o motivo formal por trás da escolha de cada papel.

---

### BPR-006 — Comunicação de status de sellers em onboarding nas dailys

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2025-08-21 |
| **Categoria** | `best_practices` |
| **Tags** | `marketing-eventos`, `audience-cs` |

**Recomendação:** Para manter todos os envolvidos alinhados durante o período de onboarding/migração em lote, o CS deve informar proativamente no grupo ou na daily:

- **Nome do cliente e etapa atual** (ex.: vínculo de anúncios, go-live, importação de catálogo).
- **Bloqueios ou impedimentos** com contexto suficiente para quem não está acompanhando.
- **Confirmação de agendas** realizadas no dia e próximas agendas previstas.
- **Clientes sem resposta** e prazo esperado para retorno.

**Rationale:** Isso permite que lideranças priorizem recursos, identifiquem escalações necessárias e evitem que clientes críticos fiquem sem visibilidade.

---

### BPR-007 — Configuração do SKU ANYMARKET para migração de sellers com presença em marketplaces

| Campo | Valor |
|---|---|
| **Grupo/Chat** | ANYMARKET & WAKE (TÁTICO / OPERACIONAL) |
| **Data** | 2025-08-28 |
| **Categoria** | `best_practices` |
| **Tags** | `integração`, `audience-solutions`, `audience-cs` |

**Recomendação:** Ao migrar um seller que já possui produtos no Magalu para o ANYMARKET, o campo **SKU ANYMARKET** deve ser preenchido com o **SKU já existente na Magalu** para garantir o vínculo correto entre os anúncios existentes e os produtos importados da Wake. Em caso de dúvida sobre qual SKU usar, verificar a planilha de cadastro no FOC (Ferramenta de Onboarding de Catálogo) do Magalu.

**Rationale:** Sem essa correspondência, o ANYMARKET cria novos anúncios onde já existem anúncios ativos, gerando duplicações. Preencher o SKU corretamente permite que o ANYMARKET gerencie os anúncios existentes em vez de duplicá-los.

---


---


### FAQ-001 — Referências de clientes de cosméticos na base Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2025-07-22 |
| **Categoria** | `faq` |
| **Tags** | `proposta`, `audience-sales` |

**Pergunta:** Quais são as referências de cosméticos na base Wake?

**Resposta compilada pelo time:**
- **Wake Experience (CRM):** Shiseido, Nars, Drunk Elephant, Mahogany (relação um pouco estremecida)
- **Wake Commerce:**
  - B2C Store (maior em GMV): https://www.b2cstore.com.br/
  - Dudah Beauty
  - LP Beauty
  - Joico: https://www.joico.com.br/
  - Doce Beleza: https://www.docebeleza.com.br/
  - Beauty Color Company: https://loja.beautycolorcompany.com.br/
  - Gel Niche: https://www.gelniche.com.br/
  - Lojas Livia: https://www.lojaslivia.com.br/
  - Dermapelle: https://www.dermapelle.com.br/
  - Cadiveu: https://store.cadiveu.com/ (atenção: havia rumores de saída no período)
  - Ybera (via B2C Store)

---

### FAQ-002 — Como indicar integradores parceiros para clientes

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2025-08-15 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-sales`, `audience-solutions` |

**Pergunta:** Como indicar integradores parceiros para um cliente?

**Resposta:**
Quando for indicar um integrador para um cliente, preencha o formulário oficial abaixo. Os Partner Ops farão a ponte com o integrador:

**Formulário:** https://docs.google.com/forms/d/e/1FAIpQLSfDs4kbyGpLUZTX18Y1Y0VTH_akMpX5Y_mHbLvB9V0toO9SMw/viewform

**Exemplo de contexto:** Para clientes que precisam de integração com Winthor (TOTVS), a Vtrina é uma opção válida. Mas sempre utilize o formulário para que os Partner Ops formalizem o processo.

---

### FAQ-003 — Marketplaces integrados com Anymarket

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2025-08-28 |
| **Categoria** | `faq` |
| **Tags** | `ecossistema-parceiros`, `audience-sales`, `audience-solutions` |

**Pergunta:** Quais marketplaces estão integrados com a Anymarket?

**Lista completa de marketplaces integrados com a Any (agosto 2025):**

Aliexpress, Amazon, Americanas, Angeloni, Angeloni (Loja Física), Armazém Paraíba, Banco Inter, Band Shop, Beleza na Web, Bradesco Shop, Camicado, Carrefour, Casa & Vídeo, Casas Bahia, Centauro, Comfortflex, Correios, CSM, CSU, Dafiti, Decathlon, Digigrow, Efácil, EOS, Época Cosméticos, Estante Virtual, Extreme, FarmaDelivery, Fast Shop, Ferreira Costa, FestaLab, GPA, Grupo Martins, Grupo Mateus, House of Gamers, IGA, Itaú, KaBuM, Koerich, Le Biscuite, Leroy Merlin, Livelo, LL Loyalty, Loja do Cooperado, Lojas Colombo, Lojas Império, Lojas Lebes, Lojas Quero Quero, Lojas MM, MadeiraMadeira, Magalu, Mercado Livre, Mobly, Netshoes, Olist, Pague Menos, Pitstop, Posthaus, Privalia, Raia, Ramarim, Renner, Ri Happy, Riachuelo, Senff Shopping, Shein, Shopee, ShopHub, Shopping BB, Sicoob, Sicredi, Steck, Sublimity, Tenda Atacado, TikTok Shop, Trocafone, Vale Bônus, VentureShop, VIK, Voce.com, Webcontinental, WeHouse, Zema, Zoom.

**Nota:** Rappi e iFood não estavam integrados com a Anymarket nesse período. Alternativas: Floui ou Iota (dependendo do contexto do projeto).

---

### FAQ-004 — Cases de e-commerce de mercado (grocery/food)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2025-09-03 |
| **Categoria** | `faq` |
| **Tags** | `proposta`, `audience-sales` |

**Pergunta:** Quais são os cases de mercado (supermercado/alimentação) na base Wake?

**Resposta (verificar atualidade antes de usar):**
- https://www.mercadinhosaudavel.com.br/
- https://www.destromacro.com.br
- https://www.lojanovamix.com.br
- https://www.saovito.com/
- https://www.ifrutus.com.br/
- https://www.lojazaeli.com.br

**Importante:** Sempre double-check se os clientes ainda estão ativos na plataforma antes de usar como referência.

---

### FAQ-005 — Cases com banner full de home com vídeo no carrossel

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2025-09-17 |
| **Categoria** | `faq` |
| **Tags** | `proposta`, `audience-sales` |

**Pergunta:** Temos algum case com banner full de home com a primeira opção do carrossel como um vídeo?

**Referências indicadas:**
- https://www.gregory.com.br
- https://www.orit.com.br/

---

### FAQ-006 — Nomenclatura correta do checkout headless Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2025-08-12 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-sales` |

**Pergunta:** Como devemos chamar o checkout headless nas conversas com clientes?

**Resposta:**
A nomenclatura oficial adotada é **"Storefront 2.0"** (não "checkout headless"). Existe um treinamento com PPT sobre o tema no SalesHub no Coda, com exemplos de quem usa e mais informações.

**Exemplo de cliente que usa Storefront 2.0:** https://www.curavino.com.br/

---

### FAQ-007 — Contato na Clear Sale para clientes

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2025-09-08 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-sales`, `audience-cs` |

**Pergunta:** Temos algum contato dentro da Clear Sale para passar para um cliente?

**Resposta:**
- **Contato:** Cauê Cruz
- **WhatsApp:** (11) 97053-7839
- **E-mail:** caue.cruz@clear.sale

---

### FAQ-008 — Clientes que migraram da Tray para Commerce Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2025-08-13 e 2025-08-25 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-sales` |

**Pergunta:** Quais clientes migraram da Tray para Wake Commerce?

**Resposta:**
- https://www.serralheria.com.br/ (migrou de Tray)
- Horizon Play: "Veio da Tray faturando 400K e já está quase em 3 milhões/mês"

**Como puxar mais dados:** Existe um relatório no Salesforce com as lojas que migramos de Tray, ordenado por MRR:
`https://locaweb.lightning.force.com/lightning/r/Report/00OU400000VPhKsMAL/view`

---

### FAQ-009 — Como acessar a conta demo do Wake Experience

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2025-07-10 |
| **Categoria** | `faq` |
| **Tags** | `proposta`, `audience-sales` |

**Pergunta:** Como acessar a conta demo do Wake Experience?

**Resposta:**
Para acessar a conta demo do Experience, a orientação oficial do time de Produto e Infra é que **cada pessoa faça a solicitação de login individual**. O login criado permite acesso a todas as plataformas.

As instruções de como solicitar o acesso estão no **Portal** da Wake. Em caso de dúvidas, contate o responsável de produto.

---

### FAQ-010 — Questionários de segurança da informação: quem envolver

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2025-07-11 |
| **Categoria** | `faq` |
| **Tags** | `integração`, `audience-sales`, `audience-solutions` |

**Pergunta:** Quando vamos responder questionários de segurança (de clientes), quem devemos envolver?

**Resposta:** O time de pré-vendas/Sales Ops deve ser acionado para coordenar a resposta de questionários de segurança da informação enviados por clientes. Em casos específicos, o time técnico/infra pode ser envolvido dependendo do escopo das perguntas.

---

## Troubleshooting

### TRB-001 — Rumor de que Wake não tem integração com Dafiti

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2025-08-12 |
| **Categoria** | `troubleshooting` |
| **Tags** | `ecossistema-parceiros`, `audience-sales`, `audience-cs` |

**Sintomas:** Um cliente (Blueman) estava dizendo para a Dafiti que a Wake não tem mais integração com eles, e o boato estava se espalhando pelo mercado.

**Causa raiz:** A informação era incorreta e havia partido da própria Dafiti para o cliente. Após uma call com a Blueman, o CS explicou que a informação não procedia.

**Resolução:**
- A integração Wake ↔ Dafiti existe e está ativa via Wake Hub/Anymarket
- Quem fica como ponto focal sobre integrações com marketplaces 3P: time de Parcerias (Henrique Viana ou Letícia)
- Para desmistificar com o cliente: reforçar que o Wake Hub (Anymarket) faz a integração com a Dafiti, e que inclusive os clientes estão sendo migrados para essa nova tecnologia

---

### TRB-002 — CS prometendo setup grátis de Experience para clientes Commerce

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2025-08-14 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-sales`, `audience-cs` |

**Sintomas:** Times de CS estavam dizendo para clientes Commerce que ao contratarem o Experience, o setup seria gratuito. Isso chegava já nas OPS como condição estabelecida, causando problemas no processo comercial. Exemplo concreto: aconteceu na proposta da Plenitude.

**Causa raiz:** Falta de orientação clara ao time de CS sobre quais concessões comerciais podem ser prometidas em upsell de Experience para clientes Commerce — o CS estava cravando setup gratuito por conta própria sem alinhamento com o time comercial.

**Resolução:**
- **Nunca confirmar setup grátis** para clientes Commerce que serão indicados para Experience
- A orientação correta é dizer que existe **"uma condição muito especial"**, mas sem cravar o valor
- Para contratos que o executivo estiver negociando com Commerce + Exp simultaneamente: **não dar o desconto no Exp sem falar com o vendedor/gerente responsável**
- Esta orientação também deve ser repassada ao time de CS

---

### TRB-003 — Problema com integração Prateleira Infinita e máquinas Pagar.me

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2026-06-03 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-sales`, `audience-solutions` |

**Sintomas:** Clientes interessados em Prateleira Infinita (Connect) assumem que as maquininhas da Pagar.me que já possuem são compatíveis com a solução — gerando ruído, frustração e atrasos no projeto quando descobrem que não são.

**Causa raiz:** Nem todas as POS da Pagar.me são compatíveis com a Prateleira Infinita. Sem uma negociação comercial específica com a Pagar.me para a operação, o cliente não recebe maquininhas adequadas.

**Resolução:**
- Todo cliente interessado em Prateleira Infinita deve passar por uma **negociação comercial específica com a Pagar.me**
- Nessa negociação, o time comercial da Pagar.me realiza:
  1. Credenciamento de uma nova conta dedicada à operação
  2. Envio das maquininhas compatíveis com o produto
- **Nunca assumir** que as maquininhas que o cliente já possui funcionarão com a Prateleira Infinita
- Esse alinhamento é fundamental para evitar: ruídos de comunicação, expectativas incorretas, atrasos nos projetos e desgaste na experiência do cliente

---

### TRB-004 — Link de assinatura com problema (Óticas Online)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2025-06-25 |
| **Categoria** | `troubleshooting` |
| **Tags** | `contrato-comercial`, `audience-sales` |

**Sintomas:** O link de assinatura do contrato da Óticas Online apresentou problema e precisou ser recriado.

**Causa raiz:** [fonte original sem informação] — a causa técnica do link defeituoso não foi registrada na conversa.

**Resolução:** O time de Sales Ops resubiu o link corrigido, e o contrato foi assinado com sucesso logo após. Lição: ao detectar problema com link de assinatura, acionar imediatamente o time de Sales Ops para recriar o link. Não esperar o cliente reportar.

---

### TRB-005 — Cliente solicita disparo/mensagens avulsas — onde está a tabela de preços

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2025-07-11 |
| **Categoria** | `troubleshooting` |
| **Tags** | `integração`, `audience-sales` |

**Sintomas:** Cliente solicita disparo/mensagens avulsas no Experience e o time comercial não encontra a tabela de preços para responder.

**Causa raiz:** A tabela de preços de disparos/mensagens avulsas não está disponível publicamente nos canais consultados pelo time comercial — só está acessível via time de Experience/produto.

**Resolução:** Consultar o gerente de Experience ou o SalesHub/Coda para versão atual da tabela de preços. Não responder ao cliente com valores sem confirmar com o time de produto.

---

## How-To

### HTO-001 — Como funciona o processo de fechamento de contratos (War Room)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2025-06-23 a 2025-06-30 |
| **Categoria** | `how_to` |
| **Tags** | `suporte-escalação`, `audience-sales` |

**O War Room de fechamento é a metodologia utilizada no sprint final do mês para garantir o maior número possível de contratos assinados.**

**Como funciona:**

1. **Sala física:** Time de Sales Ops fica disponível presencialmente na sala designada (ex: sala Credisfera, sala WAKE)
2. **Link de War Room permanente:** https://meet.google.com/hxy-iedb-tie (para quem estiver remoto)
3. **Frequência de updates:** 2x ao dia nos últimos dias do mês (manhã e final de tarde)
4. **Status Board:** Toda oportunidade é classificada em:
   - Contratos Assinados (Ganhos)
   - Em Assinatura (link enviado, aguardando assinaturas do cliente)
   - Com o Jurídico (elaborando minuta ou link)
   - Em Negociação (commit do executivo)

5. **Para cada deal em assinatura, são acompanhadas:**
   - Quais pessoas precisam assinar do lado do cliente
   - Quantas assinaturas já foram coletadas
   - Data que o link foi enviado

6. **Regras de corte do mês:** Contratos sem assinatura até o prazo de corte podem ser reconhecidos no mês seguinte mediante acordo com a gestão

---

### HTO-002 — Como registrar leads captados em eventos (Fórum ECBR, Wake Summit)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2025-08-04 a 2025-08-11 |
| **Categoria** | `how_to` |
| **Tags** | `marketing-eventos`, `audience-sales`, `audience-pam` |

**Processo para registrar contatos captados em eventos:**

**Opção 1 - Formulário:**
Preencher o formulário de captação de contatos:
`https://docs.google.com/forms/d/1wwgkvOHdTKGTSPkWxdK891cyRY1j4dBRftAX5keOGSc/edit`

**Opção 2 - Direto no Salesforce:**
Criar o lead ou contato diretamente no Salesforce, lembrando de **atribuir à campanha do evento correspondente** (ex: Fórum ECBR 2025). Em caso de dúvidas sobre como fazer isso, consultar o time de Sales Ops.

**Resultado do processo pós-Fórum ECBR 2025:**
- 32 contatos já existentes no SF foram atribuídos à campanha do pós-fórum
- 28 novos leads criados e atrelados à campanha, todos atribuídos aos respectivos proprietários do time comercial

---

### HTO-003 — Como funciona a parceria com AWS (ISV Accelerate Program / Co-sell)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2025-06-27 a 2025-07-04 |
| **Categoria** | `how_to` |
| **Tags** | `integração`, `audience-sales`, `audience-pam` |

**Contexto:** A Wake entrou no AWS ISV Accelerate Program, operando em modelo co-sell.

**Como funciona:**
1. Quando uma oportunidade sobe no Salesforce, os Account Managers da AWS ganham visibilidade e podem acionar o time Wake para apoiar
2. O inverso também acontece: AWS traz oportunidades originadas por eles (**AWS Originated**)
3. Em muitos casos, os AMs da AWS trazem recursos estratégicos como:
   - Gratuidade temporária de serviços AWS
   - Suporte técnico especializado
   - Benefícios comerciais
   - Visitas ao CD da Amazon
   - Uso do escritório da AWS para eventos

**Pontos de atenção:**
- Alguns AMs da AWS chegam mais para pedir apoio do que para apoiar — isso é natural no início da parceria
- O processo evolui gradualmente — com o tempo se tornará orgânico
- Contato principal na AWS: Iuri Nascimento (Partner Sales Manager Retail & CPG)

**Para novos leads via AWS:**
Sinalizar ao responsável de parcerias (Rafael) quais contas da carteira AWS são de interesse para que ele faça a ponte com o Account Manager correspondente da AWS.

---

### HTO-004 — Como indicar parceiros tech para concorrer à campanha de premiação

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Comercial Cross WAKE |
| **Data** | 2025-09-12 a 2025-09-16 |
| **Categoria** | `how_to` |
| **Tags** | `lead-pipeline`, `audience-pam`, `audience-sales` |

**Processo para indicar clientes a parceiros tech (e participar de corrida de prêmios):**

Preencher o formulário oficial de indicação:
`https://docs.google.com/forms/d/e/1FAIpQLSfDs4kbyGpLUZTX18Y1Y0VTH_akMpX5Y_mHbLvB9V0toO9SMw/viewform`

**Importante:** Indicações feitas por outros meios (chat, e-mail, etc.) sem o preenchimento do formulário **não contam para a campanha de premiação**. O time de Parcerias acompanha pelo formulário.

---
