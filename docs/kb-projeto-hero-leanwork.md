# KB Extract — Projeto Hero <> Leanwork <> Wake
> JID: 120363417338422570@g.us | Extracted: 2026-06-04

---

## FAQ

---

### FAQ-1 — Diferença entre "bilhete" e "apólice" no contexto Hero Seguros

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2025-11-12 |
| **Categoria** | `faq` |
| **Tags** | `hero-seguros`, `bilhete`, `apólice`, `glossário` |

**Pergunta:** O documento que o cliente recebe após contratar o seguro é chamado de apólice?

**Resposta:** Não. Na Hero Seguros, o documento emitido é chamado de **bilhete**, não de apólice. Esse é o termo correto e deve ser usado em toda a comunicação com o cliente e nas integrações técnicas.

---

### FAQ-2 — Ambiente de homologação da API Hero Seguros: necessidade de liberação de IP

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2025-09-17 |
| **Categoria** | `faq` |
| **Tags** | `api-hero`, `homologação`, `ip`, `autenticação` |

**Pergunta:** O ambiente de homologação da API Hero Seguros exige liberação de IP, assim como o de produção?

**Resposta:** Não. Em homologação não é necessária liberação de IP. A restrição de whitelist de IP se aplica apenas ao ambiente de produção.

- **URL base de homologação:** `https://api.homologacao.heroseguros.com.br`
- Auth na raiz: `POST /oauth/token`
- Demais chamadas sob: `/api/..`

---

### FAQ-3 — Credenciais de autenticação da API Hero Seguros (ambiente de homologação)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2025-09-17 |
| **Categoria** | `faq` |
| **Tags** | `api-hero`, `autenticação`, `credenciais`, `homologação`, `oauth` |

**Pergunta:** Como autenticar na API de homologação da Hero Seguros?

**Resposta:** A autenticação utiliza OAuth2 com grant type `password`. Exemplo de requisição de token para o ambiente de homologação:

```bash
curl --location --request POST 'https://api.homologacao.heroseguros.com.br/oauth/token' \
  --form 'grant_type="password"' \
  --form 'client_id="266"' \
  --form 'client_secret="EAWvWG4BTn31v5Lp4TJzGVigVLs7j1WPkN8Tyshw"' \
  --form 'username="wake@heroseguros.com.br"' \
  --form 'password="vkw96eoN4a16"' \
  --form 'scope=""'
```

> **Atenção:** Para testes de Garantia Estendida vinculados à Karcher, usar as credenciais do usuário `hero@karcher.com.br` (client_id 278), pois o usuário `wake@heroseguros.com.br` (client_id 266) é exclusivo do Parcela Garantida.

---

### FAQ-4 — Payload da API de cotação do Parcela Garantida (Prestamista)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2025-09-26 |
| **Categoria** | `faq` |
| **Tags** | `api-hero`, `parcela-garantida`, `prestamista`, `cotação`, `payload` |

**Pergunta:** Qual o payload correto para chamar a API de cotação do Parcela Garantida?

**Resposta:**

```bash
curl --location --request POST 'https://api.homologacao.heroseguros.com.br/api/prestamista/quotation' \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer <token>' \
  --data-raw '{
    "debt_amount": 4000.00,
    "installments": 10,
    "type_of_product": 5,
    "last_installment_date": "2026-08-27",
    "customer": {
        "age": 30
    }
}'
```

Campos obrigatórios:
- `debt_amount`: valor da dívida (compras + frete, **sem incluir o valor do seguro**)
- `installments`: número de parcelas
- `type_of_product`: ID do tipo de produto. O ID varia por parceiro (ex: Karcher usa `5`, não `4`)
- `customer.age`: idade do cliente

---

### FAQ-5 — Payload da API de cotação da Garantia Estendida

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2026-02-11 |
| **Categoria** | `faq` |
| **Tags** | `api-hero`, `garantia-estendida`, `cotação`, `payload`, `sku` |

**Pergunta:** Como chamar a API de cotação da Garantia Estendida?

**Resposta:** A busca do produto é feita por SKU (`external_type: "sku"`). Exemplo:

```bash
curl --location --request POST 'https://api.homologacao.heroseguros.com.br/api/extended-warranty/quotation' \
  --header 'Content-Type: application/json' \
  --header 'Authorization: Bearer <token>' \
  --data-raw '{
    "product_value": 2000.0,
    "searchProduct": {
        "external_type": "sku",
        "external_value": "30055"
    },
    "factory_warranty_end_date": "2026-12-01"
}'
```

> A `factory_warranty_end_date` é solicitada pela Hero, mas **não é usada como fator de cálculo** pela Hero Seguros. Para a Karcher, pode ser padronizada para 12 meses.

---

### FAQ-6 — Stack de infraestrutura recomendada para o middleware de integração

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2025-09-19 |
| **Categoria** | `faq` |
| **Tags** | `infraestrutura`, `middleware`, `stack`, `linux`, `banco-de-dados` |

**Pergunta:** Qual stack de infraestrutura usar para hospedar o middleware de integração Hero/Wake?

**Resposta (Hero Seguros):**
- Sistema Operacional: **100% Linux**
- Banco de dados: **MySQL** ou, preferencialmente, **Aurora**

A stack de desenvolvimento utilizada pela Leanwork no projeto é **C# .NET**.

---

### FAQ-7 — Regras de elegibilidade do cliente para o Parcela Garantida (Prestamista)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2025-09-08 |
| **Categoria** | `faq` |
| **Tags** | `parcela-garantida`, `prestamista`, `elegibilidade`, `regras-negócio` |

**Questão levantada:** A Fase 1 do projeto não implementou as regras de negócio de elegibilidade. Identificadas para a Fase 2:

- Idade mínima do cliente: **18 anos**
- Idade máxima: **69 anos** (calculada como: idade na adesão + tempo de empréstimo)
- Outras regras a detalhar na Fase 2

---

### FAQ-8 — Como testar a integração na loja Karcher (ambiente de preview)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2026-01-29 |
| **Categoria** | `faq` |
| **Tags** | `karcher`, `teste`, `preview`, `wake-storefront`, `checkout` |

**Pergunta:** Como testar a integração Hero na loja da Karcher sem impactar produção?

**Resposta:** Abrir uma aba anônima e seguir os passos nesta ordem:

1. Acessar primeiro: `https://karcher-center-altex.com.br/?preview-theme=d%2bvAEzPgMt3feu7mNxiD%2fQ%3d%3d`
2. Na mesma aba anônima, acessar: `https://karcher-center-altex.com.br/parceiros/wakegateway`

Esses dois passos ativam o cookie de parceiro e o grupo de pagamento correto com seguro.

> **Atenção:** Se a URL da barra de endereço começar com `checkout.`, o cookie do primeiro passo não foi capturado corretamente.

---

## Troubleshooting

---

### TROUBLESHOOTING-1 — API de cotação retorna `data: []` vazio sem erro

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2026-01-28 |
| **Categoria** | `troubleshooting` |
| **Tags** | `api-hero`, `prestamista`, `cotação`, `debug`, `type_of_product` |

**Sintoma:** A chamada à API `/api/prestamista/quotation` retorna HTTP 200 mas com `"data": []` e `"success": true`.

**Causa identificada:** O `type_of_product` enviado não corresponde ao parceiro correto. Cada parceiro cadastrado na Hero tem um ID de produto diferente.

**Resolução:**
- Verificar com o time Hero qual o `type_of_product` correto para o parceiro em questão.
- Para a Karcher: usar `type_of_product: 5` (e credenciais `hero@karcher.com.br`).
- O usuário `wake@heroseguros.com.br` (client_id 266) usa `type_of_product: 4` para testes genéricos.

---

### TROUBLESHOOTING-2 — Erro `QUOTATION_NO_PRICING_RULES` na cotação

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2026-02-03 |
| **Categoria** | `troubleshooting` |
| **Tags** | `api-hero`, `garantia-estendida`, `cotação`, `erro`, `pricing` |

**Sintoma:** A API retorna:
```json
{
  "success": false,
  "message": "Erro na requisição externa",
  "error": {
    "code": "QUOTATION_NO_PRICING_RULES",
    "message": "The price quoted is not allowed for this product and plan."
  }
}
```

**Causa identificada:** O produto não possui regras de precificação cadastradas, ou o `type_of_product` está errado (ex: enviando `4` em vez de `5` para a Karcher).

**Resolução:**
1. Confirmar com o time Hero o ID correto do tipo de produto para o parceiro.
2. Verificar se o produto possui SKU cadastrado e associado ao parceiro na plataforma Hero.
3. Para a Karcher, após mudança de `wake@` para `hero@karcher.com.br`, o ID mudou de 4 para 5.

---

### TROUBLESHOOTING-3 — Pedidos sem seguro sendo alterados para status "Separado" indevidamente

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2026-01-27 |
| **Categoria** | `troubleshooting` |
| **Tags** | `wake`, `status-pedido`, `gatilho`, `bug-produção`, `karcher` |

**Sintoma:** Pedidos comuns (sem seguro) na loja Karcher foram alterados automaticamente para o status "Separado" após a ativação do gatilho de integração Hero.

**Causa identificada:** O script de automação de status estava disparando para todos os pedidos pagos, sem verificar a flag de seguro contratado.

**Resolução implementada (paliativa — Wake):**
- Restringir transições de status: bloquear a mudança para "Separado" a partir de "Cancelado" e de "Enviado".

**Resolução definitiva (Leanwork):**
- Revisar o script para disparar o gatilho de "Separado" **apenas** para pedidos que contenham a flag de seguro ativo E estejam com status "Pago".

> **Lição aprendida:** Confirmar com o cliente se a loja já está em produção antes de ativar qualquer automação. Neste caso, a equipe não sabia que a Karcher já estava com pedidos reais chegando em produção.

---

### TROUBLESHOOTING-4 — `debt_amount` no bilhete incluindo erroneamente o valor do seguro

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2026-02-10 |
| **Categoria** | `troubleshooting` |
| **Tags** | `parcela-garantida`, `bilhete`, `debt_amount`, `cálculo`, `bug` |

**Sintoma:** O bilhete do Parcela Garantida apresentava valores incorretos. O `debt_amount` enviado na proposta incluía compras + frete + valor do seguro, quando deveria ser apenas compras + frete.

**Causa identificada:** O middleware estava somando o valor do seguro ao subtotal antes de enviar à API de proposta Hero (`/api/prestamista/proposal`).

**Resolução:** O `debt_amount` deve ser o valor **original da compra + frete**, sem nenhuma adição do valor do seguro. O cálculo do percentual do seguro é feito pela própria API da Hero com base no `debt_amount` informado.

---

### TROUBLESHOOTING-5 — Alíquota configurada não aparecendo no checkout

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2025-11-05 |
| **Categoria** | `troubleshooting` |
| **Tags** | `wake`, `alíquota`, `checkout`, `configuração`, `grupo-produto` |

**Sintoma:** O grupo de produtos e alíquotas foram criados no painel Wake, mas o valor do seguro não aparecia no checkout.

**Causa identificada:** Configuração incompleta do grupo de produtos ou falta de associação correta dos produtos ao grupo de alíquota.

**Resolução:** Contato direto com o suporte técnico da Wake via WhatsApp para revisão das configurações. A documentação de referência é: [Alíquotas - Wake Help Center](https://atendimento.wake.tech/hc/pt-br/articles/21406631857687-Al%C3%ADquotas)

**Resolução alternativa adotada (Fase 2 — Karcher):** Em vez de usar alíquotas, adotar a estratégia de **configuração baseada em grupo de pagamentos e cartão de crédito** como solução temporária até finalizar as configurações de alíquota.

---

### TROUBLESHOOTING-6 — Script da integração para de funcionar após ativação do ETL na Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2026-02-25 |
| **Categoria** | `troubleshooting` |
| **Tags** | `wake`, `etl`, `script`, `storefront`, `integração`, `gatilho` |

**Sintoma:** Após a ativação do ETL na Wake para a Karcher (que passou a preencher o campo `subtotal-somdesconto`), o script de integração parou de disparar: pedidos com seguro não alteravam status e não adicionavam o comentário com o link do bilhete.

**Causa identificada:** A estrutura de dados do pedido mudou com a ativação do ETL. O script estava se baseando em um campo ou estrutura que foi alterada pela nova configuração.

**Resolução (Leanwork/Proex):**
- Analisar qual campo/flag o script usava como trigger.
- Ajustar o script para reconhecer a nova estrutura de dados pós-ETL.
- Testar com pedidos com e sem seguro para validar que apenas os pedidos com seguro disparam o gatilho.

---

### TROUBLESHOOTING-7 — Pedido com PIX dispara gatilho de seguro indevidamente

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2026-03-23 |
| **Categoria** | `troubleshooting` |
| **Tags** | `parcela-garantida`, `pix`, `boleto`, `gatilho`, `bug-produção` |

**Sintoma:** Pedidos pagos via PIX (e potencialmente boleto) estavam disparando o gatilho do seguro, gerando bilhetes e e-mails indevidos para clientes que não contrataram o seguro.

**Causa raiz:** O script não estava validando o método de pagamento utilizado antes de disparar o gatilho. O seguro Parcela Garantida só é válido para pagamentos parcelados no cartão de crédito.

**Resolução (Leanwork):** Adicionar verificação do método de pagamento no script: o gatilho só deve ser disparado quando o ID do grupo de pagamento corresponder ao "Cartão com Seguro" (IDs configurados na Wake). Para a Karcher em produção, o ID correto é `16468` (GTW Cartão com Seguro). O grupo `15848` (GTW Cartão sem seguro) não deve gerar bilhete.

---

### TROUBLESHOOTING-8 — Pedido volta a status "Separado" após movimentação pela logística

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2026-02-10 |
| **Categoria** | `troubleshooting` |
| **Tags** | `wake`, `status-pedido`, `logística`, `separado`, `karcher` |

**Sintoma:** Após a logística alterar o status de um pedido de "Separado" para "Em preparação", o sistema voltava a mudar o pedido para "Separado" automaticamente.

**Causa raiz:** O script de integração não tinha controle de idempotência — disparava a mudança para "Separado" toda vez que o pedido fosse processado, sem verificar se já havia passado por esse status.

**Resolução implementada:**
1. **Wake (paliativo):** Bloquear transições de status para impedir que do status "Cancelado" ou "Enviado" se mova de volta para "Separado".
2. **Leanwork (definitivo):** Revisar o script para acionar a mudança para "Separado" apenas **uma vez**, somente quando o pedido estiver com status "Pago" pela primeira vez.

---

## How-To

---

### HOW_TO-1 — Como configurar o onboarding de um novo lojista (Middleware + Scripts + Wake)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2025-11-06 |
| **Categoria** | `how_to` |
| **Tags** | `onboarding`, `lojista`, `middleware`, `scripts`, `wake`, `deploy` |

**Visão geral do processo de implantação para um novo cliente:**

1. **Configuração do Middleware (Leanwork)**
   - Deploy do middleware na infraestrutura acordada
   - Configuração das credenciais do parceiro (client_id, secret, usuário)
   - Configuração do `type_of_product` correto para o parceiro

2. **Inclusão de Scripts (Leanwork/Proex)**
   - Incluir o script `guarantee_installment-hero.js` na loja Wake do cliente
   - Verificar se o script está apontando para as URLs corretas de bilhete

3. **Configuração da Loja Wake**
   - Criar grupos de pagamento: "GTW Cartão com Seguro" e "GTW Cartão" (sem seguro)
   - Configurar alíquotas (ou usar grupo de pagamentos como estratégia alternativa)
   - Mapear os IDs dos novos grupos de pagamento no ERP do cliente (ex: Tiny)
   - Liberar o cartão de crédito no checkout Storefront

4. **Ajustes de Front-end (Agência/Proex)**
   - Habilitar/desabilitar corretamente os grupos de pagamento no checkout
   - Aplicar o layout dos componentes Hero (box de seleção de seguro, labels, textos)

5. **Documentação de Onboarding**
   - A Leanwork elabora e envia por e-mail a documentação de onboarding para a agência responsável pela loja

> **Importante:** Deixar claro para a agência que após a configuração do backend, ela precisa ajustar o front-end.

---

### HOW_TO-2 — Como adicionar as páginas "Saiba Mais" para os produtos de seguro na Wake

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2025-09-03 |
| **Categoria** | `how_to` |
| **Tags** | `wake`, `saiba-mais`, `página-conteúdo`, `garantia-estendida`, `parcela-garantida` |

**Contexto:** No box de seleção de seguro, existe um link "Saiba mais" que deve apontar para uma página com detalhes do produto.

**Passos:**
1. Criar as páginas de conteúdo na plataforma Wake para cada produto de seguro (Garantia Estendida e Parcela Garantida)
2. O conteúdo base deve usar o documento fornecido pela Hero Seguros (ex: "Regras e Benefícios: Garantia Estendida e Parcelamento com Seguro")
3. Ajustar o HTML e CSS da página conforme identidade visual do lojista
4. Configurações globais de CSS aplicadas refletirão tanto na URL pública da página quanto no popup exibido no checkout

**Documentação Hero de referência:** Verificar versão mais atualizada do documento de definições de produtos com a Hero.

---

### HOW_TO-3 — Como realizar o deploy do Parcela Garantida em produção (checklist)

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2026-03-13 |
| **Categoria** | `how_to` |
| **Tags** | `deploy`, `produção`, `checklist`, `karcher`, `parcela-garantida` |

**Checklist validado para o go-live do Parcela Garantida na Karcher:**

**Wake (Leo):**
- [ ] Remover conector de teste e vincular conector real da Vindi
- [ ] Confirmar novo ID do grupo de pagamento gerado (ex: `16468 | GTW Cartão com Seguro`)
- [ ] Repassar novos IDs de grupos de pagamento para mapeamento no ERP (Tiny) via cliente

**Agência/Proex (Yuri/Sabrina):**
- [ ] Habilitar grupos de pagamento corretos no painel da loja
- [ ] Configurar scripts nos templates corretos
- [ ] Remover mascaras indevidas nos tokens (ex: `tcs_karch_`)
- [ ] Validar que o checkout exibe corretamente a opção de seguro

**Hero Seguros:**
- [ ] Confirmar credenciais de produção do parceiro
- [ ] Validar que o parceiro está cadastrado com o `type_of_product` correto

**Processo de virada:**
1. Primeiro subir o checkout Storefront (SF) e validar funcionamento independente
2. Somente após SF estável, ativar os scripts Hero
3. Realizar pedido teste com seguro e pedido teste sem seguro
4. Confirmar recebimento de e-mail, mudança de status e link do bilhete no comentário do pedido

---

### HOW_TO-4 — Como preparar uma loja modelo/demo para eventos e apresentações

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2025-09-05 |
| **Categoria** | `how_to` |
| **Tags** | `demo`, `evento`, `loja-modelo`, `wake-summit`, `apresentação` |

**Contexto:** Para o Wake Summit 2025, foi necessário preparar uma loja modelo (`partnerhero.fbits.store`) para demonstração dos seguros.

**Recomendações:**
1. Usar uma URL que não apareça em buscas no Google e seja desconhecida do público (sem bloqueios técnicos, mas obscura)
2. Cadastrar produtos com preços próximos da realidade para não gerar percepção negativa do custo do seguro
3. Usar os percentuais reais das tabelas de precificação fornecidas pela Hero
4. Ajustar preços dos produtos para representar valores realistas no dia do evento
5. Gravar vídeos da jornada em formato `.mkv` e compartilhar como arquivo (o WhatsApp reduz a qualidade)
6. Para apresentação em tablet: resolução recomendada de **1024px** (4:3), testado com Google Chrome

---

### HOW_TO-5 — Como mapear novos grupos de pagamento no ERP (Tiny) para a Karcher

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2026-02-23 |
| **Categoria** | `how_to` |
| **Tags** | `wake`, `erp`, `tiny`, `grupo-pagamento`, `karcher`, `mapeamento` |

**Contexto:** Ao configurar novos grupos de pagamento no checkout Wake (Storefront), os IDs precisam ser mapeados no ERP do cliente para que os pedidos sejam identificados corretamente.

**Exemplo de lista de IDs para mapeamento — Karcher:**

| ID | Descrição |
|---|---|
| 1172 | Cartão |
| 1174 | Boleto |
| 5409 | Pix |
| 12644 | Link de pagamento |
| 16307 | GTW Cartão com Seguro |
| 15848 | GTW Cartão |
| 15846 | GTW Boleto |
| 15847 | GTW Pix |
| 16468 | GTW Cartão com Seguro (produção) |

**Processo:** A solicitação de mapeamento deve ser feita pelo próprio cliente (Karcher) para o ERP (Tiny), pois a agência não tem credenciais de acesso. A Wake pode acionar via ticket se necessário.

> **Atenção:** Ao trocar de conector de teste para conector real (Vindi), um novo ID de grupo é gerado e precisa ser remapeado na Tiny.

---

## Best Practices

---

### BEST_PRACTICES-1 — Nunca fazer deploy em produção numa sexta-feira

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2026-02-20 |
| **Categoria** | `best_practices` |
| **Tags** | `deploy`, `produção`, `risco`, `boas-práticas` |

**Situação:** Surgiu a oportunidade de fazer o deploy de uma configuração em produção numa sexta-feira à tarde.

**Decisão tomada:** Não realizar o deploy na sexta-feira. Reservar uma janela na segunda-feira cedo.

**Rationale:** Deploys em produção no final da semana reduzem a capacidade de resposta a incidentes. Em caso de problema, a equipe estará menos disponível para resolver. Sempre preferir segunda-feira ou terça pela manhã.

---

### BEST_PRACTICES-2 — Confirmar o ambiente (teste vs. produção) com o cliente antes de qualquer automação

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2026-01-27 |
| **Categoria** | `best_practices` |
| **Tags** | `produção`, `teste`, `comunicação`, `onboarding`, `gestão-riscos` |

**Situação:** A equipe de desenvolvimento ativou um gatilho de automação acreditando que a loja estava apenas em ambiente de teste. Na verdade, a loja Karcher já estava em produção com pedidos reais de clientes. Resultado: pedidos reais foram afetados (status alterado indevidamente para "Separado").

**Lição:** Antes de ativar qualquer automação, script ou webhook em um ambiente de cliente, confirmar explicitamente:
1. A loja já está em produção recebendo pedidos reais?
2. Qual a URL do ambiente de testes vs. produção?
3. Existe isolamento entre os ambientes (configurações separadas)?

**Ação corretiva:** Ao ativar integrações, sempre testar com um pedido com seguro e um sem seguro para validar o escopo do gatilho.

---

### BEST_PRACTICES-3 — Estratégia de rollback: usar conector temporário para desacoplar dependências de gateway

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2026-02-04 |
| **Categoria** | `best_practices` |
| **Tags** | `vindi`, `gateway`, `rollback`, `wake`, `conector`, `testes` |

**Situação:** A Vindi (gateway de pagamento) estava em período de alta demanda (Black Friday/pós-BF) e não podia ativar novos conectores de produção. Isso bloqueava os testes completos da integração Hero.

**Solução adotada:** Usar um "conector de teste" (sem Vindi) temporariamente, mantendo o mesmo grupo de pagamento para não precisar alterar nada no script de validação. Com isso:
- Os testes do seguro (Hero) podiam prosseguir
- A única diferença: o status de "Pago" precisava ser alterado manualmente durante os testes
- Os demais passos (status "Separado", comentário com bilhete) continuavam automáticos

**Quando usar:** Sempre que uma dependência externa (gateway, ERP, etc.) bloqueie testes, avaliar se é possível desacoplar temporariamente sem comprometer a lógica de negócio.

> **Atenção:** Ao retornar para o conector real, um novo ID de grupo de pagamento é gerado e precisa ser remapeado no ERP.

---

### BEST_PRACTICES-4 — Documentar formalmente decisões técnicas tomadas durante war rooms

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2025-11-25 |
| **Categoria** | `best_practices` |
| **Tags** | `war-room`, `documentação`, `decisão-técnica`, `gestão-projeto` |

**Situação:** Após uma war room em que múltiplas equipes (Hero, Wake, Leanwork, Proex) definiram uma mudança de estratégia de alíquotas, a decisão foi resumida no grupo com os responsáveis de cada área:

> "Lean/Wake: Adotará nova estratégia para envio das porcentagens de seguro, utilizando configurações baseadas em grupo de pagamentos e cartão de crédito. Proex: atualizações necessárias no front-end. Lean: ajustes no front e back para incluir a informação vinda da URL da policy."

**Boa prática:** Ao sair de uma war room, publicar no grupo um resumo estruturado com:
- **Decisão tomada** por cada parte
- **Responsável** por cada ação
- **Prazo** ou próximo checkpoint

Isso evita retrabalho, mal-entendidos e serve como registro histórico para auditoria ou futuros projetos.

---

### BEST_PRACTICES-5 — Separação de fases para go-live: SF primeiro, seguro depois

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2026-02-27 |
| **Categoria** | `best_practices` |
| **Tags** | `deploy`, `storefront`, `go-live`, `estratégia`, `risco` |

**Contexto:** No deploy do Parcela Garantida para a Karcher, o go-live foi dividido em duas etapas:

1. **Etapa 1:** Deploy do novo checkout Wake Storefront (SF) — valida se o novo checkout funciona sem dependências do seguro
2. **Etapa 2:** Após validação do SF, ativar os scripts e a integração Hero

**Benefícios:**
- Se o novo checkout apresentar problema, não é possível imputar à integração Hero
- Janelas de rollback são mais claras para cada componente
- O cliente (Karcher) valida o checkout independentemente antes de acrescentar complexidade

**Recomendação:** Adotar esse modelo de deploy faseado em todos os projetos de implantação de seguros em lojas já existentes.

---

### BEST_PRACTICES-6 — Gestão de transição de PO durante o projeto

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2025-11-10 |
| **Categoria** | `best_practices` |
| **Tags** | `gestão-projeto`, `transição`, `po`, `continuidade` |

**Situação:** O PO responsável pelo projeto saiu para uma nova oportunidade profissional durante a execução do projeto. A transição foi gerenciada da seguinte forma:

- **Repasse gradual** nas semanas anteriores à saída
- A nova PO (Rebeca) assumiu acompanhando mais de perto o backlog de desenvolvimento
- Comunicado formal ao grupo com agradecimentos e clareza sobre continuidade
- Projeto **sem impacto no escopo ou cronograma** após a transição

**Boa prática:** Iniciar a transição de gestores de projeto com pelo menos 2-3 semanas de antecedência, com sessões de repasse incluindo o time técnico, não apenas documentação.

---

### BEST_PRACTICES-7 — Configurar e testar o produto de seguro para múltiplos SKUs antes do go-live

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2026-05-27 |
| **Categoria** | `best_practices` |
| **Tags** | `garantia-estendida`, `sku`, `produto`, `homologação`, `elegibilidade` |

**Situação:** Durante os testes finais da Garantia Estendida para a Karcher, o time ficou travado pois o produto teste inicial não estava cadastrado com um SKU elegível na plataforma Hero, e a equipe não tinha acesso ao painel da loja para cadastrar produtos.

**Lições:**
1. **Antes de iniciar testes:** definir com clareza quais SKUs estão cadastrados na Hero para o parceiro e quais produtos na loja Wake correspondem a esses SKUs
2. **Responsabilidade de cadastro:** A Wake/lojista/agência controla o painel da loja; a Hero controla a tabela de SKUs elegíveis. É necessário que ambas as listas estejam sincronizadas
3. **Lista de SKUs de teste em homologação Hero:** manter uma lista pré-cadastrada de SKUs de teste válidos para acelerar a fase de homologação de novos parceiros

**SKUs de teste utilizados em homologação:**
- `93986140`, `14287000`, `93984630`, `93984640`

---

### BEST_PRACTICES-8 — Uso da loja demo Hero para demonstrações comerciais e novos parceiros

| Campo | Valor |
|---|---|
| **Grupo/Chat** | Projeto Hero <> Leanwork <> Wake |
| **JID** | `120363417338422570@g.us` |
| **Data** | 2025-09-11 |
| **Categoria** | `best_practices` |
| **Tags** | `loja-demo`, `comercial`, `novos-parceiros`, `wake`, `hero` |

**Contexto:** Após a apresentação no Wake Summit com a loja modelo `partnerhero.fbits.store`, a Hero fechou o primeiro contrato de parceiro (Karcher) usando exatamente essa demonstração.

**Padrão de sucesso identificado:**
- Ter uma loja modelo funcional e com dados realistas acelera o ciclo de vendas
- A demonstração ao vivo de toda a jornada (seleção de GE + Parcela Garantida no checkout) foi mais efetiva que slides
- Gravar vídeos da jornada para uso em eventos onde não é possível demonstrar ao vivo

**Recomendação para próximos parceiros:** Manter a loja demo atualizada com a versão mais recente da integração e com preços de produtos realistas, para que a relação custo-benefício do seguro seja percebida positivamente pelo cliente em potencial.
