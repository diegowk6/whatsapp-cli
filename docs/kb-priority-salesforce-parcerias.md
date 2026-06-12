# KB Prioritário — Salesforce para Operações de Parcerias

**Tópico**: Manual canônico do Salesforce (SF) para operadores de parcerias e canais — registro de leads, carrossel de distribuição, tipos de indicação, ciclo de OPT, acesso ao Community e bugs conhecidos.

**Audiências**: `audience-pam` (primária), `audience-sales` (secundária).

**Fontes primárias**: `WAKE | Time Agência` e `WAKE | Time Parcerias` (mensagens 2025-06 → 2026-05) + extrações em `docs/kb-delivery-batch-01.md` (FAQ-001 a FAQ-012, TRB-001 a TRB-005, HTO-001 a HTO-003 do bloco kb-agencia).

**Versão**: 1.0 — 2026-06

---

## Como usar este documento

Este é um manual operacional, não uma referência completa do schema do SF. Cada seção responde a uma pergunta concreta que aparece com frequência nos grupos do time de Parcerias e Agências. Toda afirmação sobre comportamento do SF está marcada com:

- ✓ verificado — confirmado por mensagem citada em `store/messages.db`
- ⚠ inferido — extraído de prática operacional repetida, sem mensagem canônica única
- ❌ sem fonte — não substanciado (não deve aparecer no corpo final; presente só se a equipe pediu nota de gap)

A regra mais recente prevalece. Quando o comportamento mudou ao longo do tempo, a versão anterior é marcada como histórica com data.

---

## 1. Como registrar um lead vindo de parceiro

### 1.1 Decisão inicial: qual é a origem real do lead?

Antes de tocar em qualquer campo, decidir entre três categorias mutuamente exclusivas. Esta é a taxonomia canônica usada pelo time desde 2025-07-11 (✓ verificado, `WAKE | Time Agência`, 2025-07-11):

| Tipo de indicação | Quando usar | Efeito no carrossel |
|---|---|---|
| **Venda de Canal** | Lead aberto e operado por executivo Wake interno (Comercial ou Canais). Parceiro/agência não originou. | Carrossel pode rodar se "Parceiro Operador" estiver vazio. |
| **Indicação Agências** | Lead originado por uma agência específica (ex.: GE, Catus, Eficaz indicou explicitamente). | Carrossel **não deve** rodar — agência indicadora vai em "Parceiro Indicado 1" e a indicação é direta. |
| **Indicação Parcerias** | Lead originado por parceiro de tecnologia/martech (ex.: Bonifiq, CRMBonus, Vurdere) que indicou para Wake. | Carrossel pode rodar para distribuir entre agências, **se** a agência ainda não foi escolhida. |

✓ verificado — `WAKE | Time Agência`, 2025-07-11 (lista textual: "Venda de Canal / Indicação agências / Indicação parcerias").

### 1.2 Field-by-field: o que preencher

Os campos críticos para um lead/OPT de parceria ficam na **aba Parcerias** da Oportunidade. Em julho/2025 o campo `Parceiro Operador` foi sugerido para mover da aba Detalhes para a aba Parcerias justamente para reduzir esquecimento de preenchimento (⚠ inferido sobre o status atual da mudança — sugestão registrada em 2025-07-24, sem confirmação posterior de deploy).

| Campo | O que significa | Quando preencher |
|---|---|---|
| **Origem do Lead** | De onde a oportunidade veio (Parcerias e Canais, Marketing, Inbound, Indicação Interna, Venda de Canal etc.). | Sempre. Define o relatório que conta a indicação. |
| **Solução de interesse** | Wake Commerce, Wake Experience, Storefront, OMS, Creators, DXP. | Sempre. **Só Wake Commerce dispara carrossel** (✓ verificado, 2026-02-10). |
| **Tipo da Oportunidade** | Novo contrato, Renovação, Upsell. | Sempre. **Só "Novo contrato" dispara carrossel** (✓ verificado, 2026-02-10). |
| **Divisão** | Comercial, Canais, Enterprise, Corp. | Sempre. **Apenas Comercial ou Canais disparam carrossel** (✓ verificado, 2026-02-10). |
| **Parceiro Operador** | Agência que vai operar a implementação. Quando preenchido, **trava o carrossel** — a OPT é direcionada manualmente. | Preencher quando a agência já está definida (indicação direta, summit, evento patrocinado, projeto manual). |
| **Parceiro Indicado 1, 2, 3** | Agências indicadas pelo carrossel ou inseridas manualmente para terem visibilidade da oportunidade. Não há diferenciação de peso entre 1, 2 e 3 (✓ verificado, 2025-12-16). | Carrossel preenche automaticamente. Em indicação manual, preencher a aba Parcerias retirando "Parceiro Operador" e adicionando as agências escolhidas. |

### 1.3 Regra prática: Parceiro Operador vs Parceiro Indicado

- **Parceiro Operador preenchido** → uma agência única e definida vai operar a conta. Carrossel não roda. Use para: indicações diretas de uma agência específica, summits onde a agência patrocinou junto, projetos manuais com agência pré-acordada.
- **Parceiro Indicado 1/2/3 preenchidos (com Operador vazio ou "A Definir")** → 3 agências receberão a oportunidade e o cliente escolhe. Use para: distribuição via carrossel (preenchimento automático), ou indicação manual de 3 agências quando o time decidiu não usar o carrossel.

✓ verificado — `WAKE | Time Agência`, 2025-07-24 (instrução textual: "Precisamos retirar o 'Parceiro Operador' e incluir em 'Parceiro Indicado 1' e 'Parceiro Indicado 2' a Wicomm e a Oasis").

✓ verificado — `WAKE | Time Agência`, 2026-02-06 (Série//A, no contexto de eventos patrocinados: "definimos o parceiro operador que patrocinou conosco, e não acarreta em rodar o carrossel de distribuição").

---

## 2. Como o carrossel funciona

### 2.1 Algoritmo de disparo (regras simultâneas)

O carrossel **roda automaticamente** apenas quando **todas** estas quatro condições são verdadeiras na Oportunidade:

1. **Solução de interesse** = `Wake Commerce`
2. **Parceiro Operador** = `nulo` (vazio)
3. **Tipo da Oportunidade** = `Novo contrato`
4. **Divisão** = `Comercial` ou `Canais`

Se qualquer condição falhar, o carrossel não dispara. As três condições adicionais (Solução, Tipo, Divisão) também são usadas para direcionar a OPT no atributo Enterprise/Canais/Corp.

✓ verificado — `WAKE | Time Agência`, 2026-02-10 (mensagem da Renata para Aninha, com a lista numerada exata dos 4 campos).

### 2.2 Quando exatamente o carrossel dispara dentro do ciclo

O disparo ocorre **na etapa de qualificação da OPT**, automaticamente, assim que a OPT é aberta com os 4 campos qualificados:

> "mas hoje quando abre opt, em qualificação, roda o carrossel automaticamente"
> ✓ verificado — `WAKE | Time Parcerias`, 2025-07-07.

Isso significa: o carrossel **não espera** o avanço da OPT para "Negociação" ou "Proposta". Ele atua no momento da abertura/qualificação. Por isso converter um Lead em OPT sem o `Parceiro Operador` preenchido já é suficiente para distribuir para 3 agências (ver bug em §6).

### 2.3 Lógica de distribuição entre agências

A distribuição leva em conta:
- **Nível/tier da agência** (Essential, Special, Elite — pesos diferentes, ✓ verificado por referência repetida a "pesos" em 2025-12-18 e 2025-12-16).
- **Histórico recente de indicações** — agências que já apareceram em OPTs anteriores entram em fila menor temporariamente; a indicação "reseta" quando o ciclo é refeito ou via ticket manual (✓ verificado, `WAKE | Time Agência`, 2025-07-04: "As indicações de agencias 'resetaram' entao eles voltaram a ter vez no proximo carrossel").
- **Status do carrossel da agência** — agências podem ser **pausadas** manualmente (por inadimplência, pedido de cancelamento de contrato, problemas operacionais). Pausa é feita por ticket ao time de Salesforce/Sales Ops (✓ verificado, várias mensagens 2025-12-08 a 2026-04-06).

A regra prática que sai do corpus: **3 agências por carrossel rodado** é o padrão (✓ verificado, 2025-07-11 "rodou Carrossel para Catus, Eficaz, Econverse").

### 2.4 Quando o carrossel **não** dispara

| Cenário | Por quê |
|---|---|
| Renovação de contrato | Tipo da Oportunidade ≠ Novo contrato (✓ verificado, 2026-02-10 + fix de 2025-07-04). |
| Parceiro Operador preenchido (qualquer valor real) | Trava por design (✓ verificado, 2026-02-10). |
| Parceiro Operador = "A Definir" | Trava o carrossel; usado intencionalmente pelo time Comercial quando ainda não decidiu agência (✓ verificado, 2025-12-16 e 2025-12-10). |
| Solução ≠ Wake Commerce (ex.: Storefront, OMS, Creators, DXP) | Não dispara (✓ verificado, 2026-02-10). Cross-ref com FAQ-012 do batch-01. |
| Divisão = Enterprise ou Corp | Não dispara (✓ verificado, 2026-02-10). |
| Conta teste vinculada como Parceiro Operador | Workaround intencional para bloquear (✓ verificado, 2025-07-03: "Podemos ter uma conta Teste de agencia que vinculamos como Parceiro operador somente no intuito de bloquear disparos de carrossel"). |

---

## 3. Tipos de indicação no SF — comportamento de cada um

Recapitulando a taxonomia da §1, com o efeito operacional de cada categoria:

### 3.1 Venda de Canal

Lead originado e operado internamente por executivo Wake (Comercial ou Canais). Sem participação de agência ou parceiro tech.
- **Origem do Lead** = Indicação Interna ou Venda de Canal
- **Parceiro Operador** = vazio (carrossel roda) **ou** agência definida posteriormente
- **Comportamento esperado**: se a Wake quer envolver agências, deixa Operador vazio e o carrossel distribui. Se já há agência alinhada (ex.: cliente pediu uma específica), preencher Operador.

### 3.2 Indicação Agências

Lead **originado por uma agência específica** que trouxe o cliente para a Wake.
- **Origem do Lead** = Parcerias e Canais (ou rótulo equivalente)
- **Parceiro Operador** = a agência que indicou (deve ser preenchido para travar carrossel)
- **Comportamento esperado**: carrossel **não roda**. A indicação é direta para a agência indicadora.

⚠ inferido — quando a agência indicadora é colocada apenas em "Parceiro Indicado 1" sem ser Operador, o sistema histórico ainda rodava carrossel para outras (caso Sapataria Nova, 2025-07-11: lead da Séri.e converteu em OPT e carrossel direcionou para Catus, Eficaz e Econverse). Ver §6 — bug TRB-003.

### 3.3 Indicação Parcerias

Lead originado por parceiro de tecnologia/martech (não-agência) — ex.: Bonifiq, CRMBonus, Vurdere indicou cliente.
- **Origem do Lead** = Parcerias e Canais
- **Parceiro Operador** = vazio (parceiro tech não opera implementação)
- **Comportamento esperado**: carrossel **roda** para distribuir entre agências, porque o parceiro tech não vai implementar.

✓ verificado — `WAKE | Time Agência`, 2025-07-07: "Goldko / Lead veio por parcerias (Bonifiq) mas não houve resposta do cliente (fique a vontade pra girar). Girei Carrossel."

---

## 4. OPT lifecycle — quando o carrossel age em cada etapa

O ciclo padrão de uma Oportunidade Wake:

```
Lead (registro inicial)
   ↓ conversão
OPT — Qualificação    ← CARROSSEL DISPARA AQUI (se condições §2.1)
   ↓
OPT — Negociação      ← carrossel já agiu; alterações exigem ticket manual
   ↓
OPT — Proposta
   ↓
OPT — Fechamento (Ganho ou Perdido)
```

### 4.1 Comportamento por etapa

| Etapa | Carrossel age? | Observação |
|---|---|---|
| Lead (pré-conversão) | Não | Indicações se ajustam editando o Lead direto. |
| Conversão Lead → OPT | **Sim, automático na qualificação** | ✓ verificado, 2025-07-07 ("quando abre opt, em qualificação, roda o carrossel automaticamente"). |
| OPT em Negociação | Não dispara novamente | Mudanças manuais (remover agência indicada errada, trocar Operador) exigem ticket Sales Ops. |
| OPT em Proposta/Fechamento | Não dispara | Indicações ficam congeladas para fins de contagem de Revenue Share. |
| OPT desqualificada e reaberta | **Pode disparar de novo** | ✓ verificado, 2025-09-29 ("opt desqualificada e reaberta") — risco de redistribuição. |
| OPT convertida com Lead já tinha Operador "A Definir" | **Não roda na conversão** | ✓ verificado, 2025-12-10: "Ele não rodou, pois foi incluído um parceiro operador como 'A Definir'". |

### 4.2 Cuidado especial: conversão de Lead com agência indicadora

Quando uma agência traz o lead, a sequência segura é:

1. Antes de converter Lead → OPT, garantir que `Parceiro Operador` = nome da agência indicadora.
2. Só então converter.
3. Se converter primeiro e preencher depois, o carrossel **já rodou** e existem 3 agências indicadas — exige ticket manual para limpar.

✓ verificado — `WAKE | Time Agência`, 2025-07-11: "Converti lead em OPT vindo da Séri.e, fiz todo o processo certo, no momento da conversão o carrossel direcionou para Catus, Eficaz, Econverse." Esta mensagem documenta que mesmo seguindo o processo, em julho/2025 a conversão disparou o carrossel — tratado como bug (§6, TRB-003).

---

## 5. Acesso ao Community Salesforce para parceiros

O Community é o portal SF onde a agência visualiza as OPTs em que foi indicada, sem precisar de licença interna do SF. É um produto licenciado: em 2026-01-16 o time avaliou reduzir o pacote de 60 licenças porque ~30 estavam em uso e quase metade sem acesso há 3 meses (✓ verificado, `WAKE | Time Agência`, 2026-01-16).

### 5.1 Como liberar acesso ao Community

Não há mensagem que documente a sequência completa de provisionamento passo-a-passo. O que está confirmado no corpus:

1. **Responsável atual**: o time de Partner Ops (Renato a partir de 2025-07-14) cuida de demandas operacionais — troca de e-mails, reset de senha, problemas com o Community (✓ verificado, 2025-07-14).
2. **Cadastro de novo usuário**: feito manualmente no SF Community pela PAM responsável (ex.: 2025-10-09 "realizei todo o cadastro no SalesForce Community para a Driven.CX").
3. **E-mail de ativação**: usuário recebe e-mail automático de alteração de senha para primeiro acesso (✓ verificado, `Duo e Wake Parcerias`, 2026-01-15).
4. **Vínculo Master vs vínculo de visualização**: agência pode ter um "Usuário Master" e usuários adicionais apenas com permissão de visualizar OPTs/Leads (⚠ inferido, `Simples Inovação & Wake Parcerias`, 2025-08-26 — questão levantada, resposta indica "creio que só teríamos como deixar uma pessoa vinculada ao Community", sem confirmação definitiva).

Para casos não cobertos: abrir ticket interno para o time de Salesforce via Renato (Partner Ops) ou via Aninha (responsável por incidentes SF/Community a partir de 2025-07-11, ✓ verificado).

### 5.2 Troubleshooting — Community vazio ou leads não aparecem

**Sintoma**: agência abre Community e não vê uma OPT que foi indicada para ela.

**Causas raiz documentadas no corpus**:

1. **Erro de cadastro original** — OPT aparecia em ambos os lados (SF interno e Community) ou em nenhum por falha de trava (✓ verificado, 2025-07-24: "Erro de Oportunidades aparecerem de ambos os lados (SF e Community): corrigido parcialmente"). Lucas (time SF) mapeou e corrigiu parte em jul/2025.
2. **Origem do lead errada** — se o lead foi cadastrado com Origem ≠ "Parcerias e Canais" mas a agência apareceu na OPT, ele não conta no dashboard de leads do parceiro nem aparece no Community filtrado (✓ verificado, 2025-08-15: "Catus abriu pelo Community e não consta no Dash... colocou origem parcerias e canais, se for isso mesmo, temos muitos leads ainda que podem não estar nessa contagem").
3. **Filtros de lista incorretos** — em jul/2025 havia filtros desnecessários na visão da agência que escondiam OPTs; foram removidos por Lucas (✓ verificado, 2025-07-24).
4. **Botão sumiu / mudança de UI** — incidente pontual em 2026-03-02 ("fizemos alguma alteração recente no community? O botão sumiu") — sem causa raiz documentada na conversa, escalado.

**Resolução padrão**: abrir ticket marcando Renato (Partner Ops) com link da OPT e print do que a agência vê. Caso seja erro de cadastro, Lucas (time SF) corrige.

### 5.3 Política de habilitação de leads condicionada a certificações

Em 2025-10-09 ficou em aberto se Community só seria liberado para a agência **após** conclusão das certificações mínimas do plano (✓ verificado, 2025-10-09 — pergunta levantada, sem mensagem posterior de resposta no corpus consultado). ⚠ inferido como regra ainda em discussão até a data de corte.

---

## 6. Bugs conhecidos e workarounds

Cada bug abaixo é citado com a data do incidente. **Quando há fix posterior, está marcado como "histórico" com a data do fix**. Sem marcação histórica = ainda recorrente até o último registro do corpus.

### 6.1 BUG-1: Carrossel disparando em renovações (✓ histórico, fix 2025-07-04)

- **Sintoma**: Samires (executiva Wake responsável por renovações) abria OPTs com prefixo "Renovação - [Cliente] - [Mês/Ano]" e o carrossel disparava, distribuindo a renovação para 3 agências que não eram a agência que já cuidava da conta. Documentado em 2025-07-03 com 2 casos concretos (Madel, Casa Bergan).
- **Causa raiz**: faltava trava no SF que verificasse `Tipo da Oportunidade = Renovação` antes de disparar carrossel. O campo "Tipo da Oportunidade" não estava sendo respeitado pelo automatismo no momento do bug.
- **Impacto**: agência incumbente (ex.: FG cuidando da Madel) podia ser preterida por outra agência que recebeu indicação errada.
- **Fix**: 2025-07-04, Sales Ops confirmou — "Tema do Carrossel resolvido e OPTs de renovação segundo o time de Sales Ops nao rodarão mais o carrossel". As indicações já feitas foram "resetadas" — agências voltaram a ter vez no próximo carrossel.
- **Estado atual (2026-02-10)**: a regra "Tipo = Novo contrato" está formalizada como condição obrigatória de disparo (✓ verificado). Bug não deveria mais ocorrer.

✓ verificado — `WAKE | Time Agência`, 2025-07-03 (relatos), 2025-07-04 (fix), 2026-02-10 (regra consolidada). Cross-ref: batch-01 TRB-001.

### 6.2 BUG-2: Agências novas recebendo mais leads que as antigas

- **Sintoma**: agências recém-onboarded apareciam em mais OPTs no período inicial do que agências antigas do mesmo tier.
- **Causa raiz**: pesos do carrossel desbalanceados favorecendo agências sem histórico recente (mecanismo de fila — quem foi indicado menos sobe na fila).
- **Workaround**: ajuste manual de pesos pelo time SF. Em 2025-12-18 o time confirmou: "iríamos mexer diretamente nos pesos e acabar desbalanceando o Carrossel". Em 2025-12-16 foi reportado como "resolvido já, referente aos pesos".
- **Estado atual**: ⚠ inferido — o time de Agência continuou registrando preocupação em 2026-04-06 ("Carrossel - possível alteração na etapa de rodagem para as agências para diminuição de distribuição erradas"), sinalizando que ajustes seguem em andamento.

✓ verificado — `WAKE | Time Agência`, 2025-12-16, 2025-12-18, 2026-04-06. Cross-ref: batch-01 TRB-002.

### 6.3 BUG-3: OPT disparou carrossel para agência errada / agência indicadora

- **Sintoma**: executivo Wake abre OPT a partir de Lead indicado por agência X, mas o carrossel dispara e indica X + outras 2 agências, ou pior, indica 3 agências e X **não** entra.
- **Causa raiz**: `Parceiro Operador` não estava preenchido no momento da conversão Lead → OPT. O carrossel age na qualificação e ignora o histórico de quem indicou o Lead original.
- **Caso documentado**: Sapataria Nova (2025-07-11) — Lead da Séri.e, conversão em OPT direcionou para Catus, Eficaz, Econverse.
- **Workaround imediato**:
  1. Abrir ticket Sales Ops com link da OPT.
  2. Solicitar remoção das agências indicadas erradas dos campos `Parceiro Indicado 1/2/3`.
  3. Preencher `Parceiro Operador` com a agência indicadora original.
  4. Solicitar que as agências indicadas erradamente sejam removidas do **relatório de contagem de indicações** (Dashboard `01ZU4000000QtP3MAK`) — caso contrário a contagem ficará distorcida e impactará o carrossel futuro.
- **Prevenção**: antes de converter Lead → OPT, sempre preencher Parceiro Operador.

✓ verificado — `WAKE | Time Agência`, 2025-07-11, 2025-07-22 ("Lembrando que essas devem ser retiradas o parceiro via carrossel e deixar apenas o origem/operador"). Cross-ref: batch-01 TRB-003, TRB-005.

### 6.4 BUG-4: Exec abrindo OPT indicada por agência e rodando carrossel automaticamente

- **Sintoma**: variante do BUG-3 — exec Wake abre OPT que já é claramente indicação de uma agência, mas esquece de preencher Operador, dispara carrossel para outras 3 agências incluindo a indicadora.
- **Causa raiz**: falha de processo do executivo + falta de trava no SF que detecte que o Lead original tinha agência indicadora.
- **Workaround estrutural** (2025-07-03):
  - Criar conta-teste de agência ("á definir" / "teste") para usar como Parceiro Operador quando o exec ainda não sabe qual agência usar — trava o carrossel sem perder vez.
  - Alinhamento com Samires (e demais execs) sobre o procedimento.
- **Mudança de UI sugerida**: mover campo `Parceiro Operador` da aba Detalhes para a aba Parcerias para reduzir esquecimento (sugerido em 2025-07-24, ⚠ inferido sobre o status atual da mudança).

✓ verificado — `WAKE | Time Agência`, 2025-07-03, 2025-07-04, 2025-07-24.

### 6.5 BUG-5: OPT sem Parceiro Operador rodou carrossel sem permissão

- **Sintoma**: PAM percebe OPT que rodou carrossel mas a conta deveria ser tratada manualmente (ex.: Eucatex, ABC da Construção via Summit, Aste).
- **Causa raiz**: convergente com BUG-3/BUG-4 — Operador vazio quando deveria estar preenchido.
- **Workaround**: ticket para remover do carrossel e fazer indicação manual via aba Parcerias (`Parceiro Indicado 1/2/3` preenchidos, `Parceiro Operador` removido ou "A Definir"). Ver §7.

✓ verificado — `WAKE | Time Agência`, 2025-07-17, 2025-07-24, 2025-10-27.

### 6.6 BUG-6: Lead da agência aberto via Community não conta no dashboard

- **Sintoma**: agência (ex.: Catus) abriu lead diretamente pelo Community, lead aparece no SF, mas não conta no Dashboard QuickSight de leads por parceiro.
- **Causa raiz**: campo Origem do Lead foi preenchido como "Parcerias e Canais" em vez do valor que o dashboard filtra.
- **Workaround**: revisar regra de origem do dashboard para incluir variações; até lá, contagem manual via relatório SF.

✓ verificado — `WAKE | Time Agência`, 2025-08-15.

---

## 7. Quando e como fazer indicação manual (sem carrossel)

Use indicação manual quando:
- Cliente pediu agências específicas (não vale rodar fila).
- Evento/Summit onde agências patrocinaram juntas com a Wake.
- Casos especiais escalados via CS Wake (ex.: ABC da Construção em 2025-10-27).
- Renovação onde uma agência específica já cuida da conta.
- Workaround para BUGs §6.

### 7.1 Procedimento — indicação manual via aba Parcerias

1. **Abrir a OPT no SF** (link `locaweb.lightning.force.com/lightning/r/Opportunity/...`).
2. **Navegar para a aba Parcerias** (não a aba Detalhes — desde a sugestão de mudança em 2025-07-24, o campo Parceiro Operador também migrou para esta aba; ⚠ inferido sobre confirmação de deploy).
3. **Remover Parceiro Operador** se estiver preenchido com valor real. Substituir por "A Definir" para travar carrossel sem indicar agência real, **ou** deixar vazio (em casos onde a OPT já passou da qualificação e o carrossel não disparará mais).
4. **Preencher Parceiro Indicado 1, 2 e 3** com as agências escolhidas. Não há diferença de peso entre os três slots (✓ verificado, 2025-12-16).
5. **Confirmar com Renato (Partner Ops)** se a alteração não conseguir ser feita via Community ou se a OPT estiver travada para edição da PAM — abrir ticket marcando Renato.
6. **Comunicar as agências indicadas** pelos grupos de WhatsApp respectivos (mesmo que Community esteja correto — comunicação ativa evita perda de OPT por inatividade).

✓ verificado — `WAKE | Time Agência`, 2025-07-24 (caso Aste com procedimento textual: "Parceiro Indicado 1: Wicomm / Parceiro Indicado 2: Oasis / Parceiro Indicado 3: FG / Precisa retirar: Parceiro operador: á definir"). Cross-ref: batch-01 HTO-001.

### 7.2 Observações

- **Não é possível** fazer indicação manual sem ticket quando o campo está travado pela visualização padrão da PAM — ⚠ inferido em algumas mensagens (2025-07-24: "não sei como fazer isso manualmente").
- Para agências em pausa (inadimplência, cancelamento), a indicação manual também precisa respeitar o status — ver §2.3.

---

## 8. Como resetar uma indicação errada

Quando uma agência foi indicada erroneamente (BUG-3, BUG-4, BUG-5 ou erro humano), o reset envolve dois passos paralelos: limpar a OPT e limpar a contagem.

### 8.1 Procedimento

1. **Identificar todas as agências indicadas erradas** na OPT (campos Parceiro Indicado 1/2/3).
2. **Abrir ticket Sales Ops** com:
   - Link da OPT
   - Lista nominal das agências a remover
   - Justificativa (ex.: "lead originado pela Séri.e, carrossel direcionou erroneamente para Catus, Eficaz, Econverse")
   - Pedido explícito de remover também do **relatório de contagem de indicações** (Dashboard `01ZU4000000QtP3MAK`, ✓ verificado em 2025-07-22 e 2025-07-23) — sem isso, as agências removidas perdem vez no próximo carrossel injustamente.
3. **Aguardar confirmação de Sales Ops** — sem SLA formal documentado (⚠ inferido — não há mensagem que estabeleça prazo).
4. **Comunicar agências afetadas**: agências removidas devem ser informadas de que não atuarão na OPT; agência correta (indicadora real) deve ser comunicada de que assumirá.
5. **Validar contagem no Dashboard** após a remoção. Em 2025-07-22 o time documentou: "essas devem ser retiradas o parceiro via carrossel e deixar apenas o origem/operador".

✓ verificado — `WAKE | Time Agência`, 2025-07-03 a 2025-07-04 (precedente do reset Samires), 2025-07-22 (procedimento), 2026-05-04 (caso Corello/Zema: "Abri ticket de ambas para remover do carrossel").

### 8.2 Quando não resetar

- Se a OPT já foi para Negociação avançada e cliente já está conversando com uma das agências indicadas → consultar Coordenação Parcerias antes (caso Dumond, `Parcerias Dotkom & Wake`, 2026-01-27: agências do carrossel já foram envolvidas, decisão do cliente prevaleceu).

---

## Fontes

### Mensagens WhatsApp citadas

- `WAKE | Time Agência`, 2025-06-27 — discussões iniciais sobre carrossel vs venda de canal.
- `WAKE | Time Agência`, 2025-07-03 — relato do bug Samires (renovações), workaround conta-teste como Parceiro Operador.
- `WAKE | Time Agência`, 2025-07-04 — fix do bug renovações, reset de indicações.
- `WAKE | Time Agência`, 2025-07-07 — taxonomia "Samires - Desconsiderar / OMS - Não vai para agência".
- `WAKE | Time Parcerias`, 2025-07-07 — "quando abre opt, em qualificação, roda o carrossel automaticamente".
- `WAKE | Time Agência`, 2025-07-11 — taxonomia textual "Venda de Canal / Indicação agências / Indicação parcerias" + caso Sapataria Nova (BUG-3).
- `WAKE | Time Agência`, 2025-07-14 — Renato (Partner Ops) assume demandas operacionais Community.
- `WAKE | Time Agência`, 2025-07-17 — pedido de ajuda em carrossel sem Parceiro Operador.
- `WAKE | Time Agência`, 2025-07-22, 2025-07-23 — procedimento de remover do carrossel + Dashboard de contagem.
- `WAKE | Time Agência`, 2025-07-24 — sugestão mover Parceiro Operador para aba Parcerias + caso Aste com indicação manual + reunião com Lucas (correções Community).
- `WAKE | Time Agência`, 2025-08-15 — bug dashboard não contabiliza lead aberto via Community (BUG-6).
- `Simples Inovação & Wake Parcerias`, 2025-08-26 — discussão sobre Master vs visualização no Community.
- `WAKE | Time Agência`, 2025-10-09 — cadastro Community Driven.CX + pergunta sobre certificações como pré-requisito.
- `WAKE | Time Agência`, 2025-10-27 — caso ABC da Construção (indicação manual via aba Parcerias).
- `WAKE | Time Agência`, 2025-12-08 — discussão sobre mudança de pesos do carrossel.
- `WAKE | Time Agência`, 2025-12-10 — carrossel não rodou porque Operador era "A Definir".
- `WAKE | Time Agência`, 2025-12-16 — Dashboard alerta OPTs com Operador "A Definir"; sem diferenciação de peso entre Indicado 1/2/3.
- `WAKE | Time Agência`, 2025-12-18 — calibragem dos pesos consultando time SF.
- `WAKE | Time Agência`, 2026-01-16 — discussão de licenças Community (60 licenças, ~30 ativas).
- `WAKE | Time Agência`, 2026-02-10 — **regra canônica das 4 condições do carrossel** (Renata → Aninha).
- `Wake <> Série//A`, 2026-02-06 — evento patrocinado e Parceiro Operador definido evitam carrossel.
- `Wake <> Série//A`, 2026-03-25 — erro de Operador preenchido mas carrossel rodou (variante BUG-3 recorrente).
- `WAKE | Time Agência`, 2026-03-02 — botão sumiu no Community (incidente UI).
- `WAKE | Time Parcerias`, 2026-04-06 — carrossel: possível alteração na etapa de rodagem.
- `WAKE | Time Parcerias`, 2026-05-04 — caso Corello/Zema removidos do carrossel.
- `WAKE | Time Parcerias`, 2026-05-28 — procedimento padrão para voltar agência ao carrossel via ticket.

### Artigos extraídos referenciados

- `docs/kb-delivery-batch-01.md` — bloco kb-agencia:
  - FAQ-001 (atuação conjunta agência+exec), FAQ-002 (carrossel em renovação), FAQ-003 (modelo de indicação que evita carrossel), FAQ-009 (lead via formulário com agência indicadora), FAQ-012 (carrossel só para Wake Commerce).
  - TRB-001 (carrossel em renovações), TRB-002 (agências novas recebendo mais leads), TRB-003 (Lead → OPT disparou carrossel errado), TRB-005 (agência indicada por conflito de interesse).
  - HTO-001 (indicação manual sem carrossel), HTO-002 (registro de lead de agência), HTO-003 (retirar agências indicadas erroneamente).
- `docs/kb-delivery-batch-03.md` — bloco kb-parcerias (referência cruzada para origem do lead e fluxo de parceiros tech).

---

## Verificação

| Claim | Status | Query / Fonte | Evidence |
|---|---|---|---|
| Carrossel dispara apenas com 4 condições (Solução=Wake Commerce, Operador=nulo, Tipo=Novo contrato, Divisão=Comercial\|Canais) | ✓ verificado | `LIKE '%carrossel%'` em `WAKE \| Time Agência` (40 hits) | WAKE \| Time Agência, 2026-02-10 (lista textual numerada) |
| Carrossel dispara na qualificação da OPT (não depois) | ✓ verificado | `LIKE '%qualifica%' AND LIKE '%carrossel%'` | WAKE \| Time Parcerias, 2025-07-07 |
| Bug de renovação disparando carrossel — fix em 2025-07-04 | ✓ verificado (histórico) | `LIKE '%Samires%' OR LIKE '%renovaç%'` (40 hits) | WAKE \| Time Agência, 2025-07-03 (bug) + 2025-07-04 (fix) + 2026-02-10 (regra consolidada) |
| Taxonomia de 3 tipos de indicação (Venda de Canal / Agências / Parcerias) | ✓ verificado | `LIKE '%Venda de Canal%'` | WAKE \| Time Agência, 2025-07-11 (lista textual) |
| Operador "A Definir" trava carrossel intencionalmente | ✓ verificado | `LIKE '%A Definir%' OR LIKE '%á definir%'` | WAKE \| Time Agência, 2025-12-10, 2025-12-16; 2025-07-24 |
| Não há diferença de peso entre Parceiro Indicado 1, 2 e 3 | ✓ verificado | `LIKE '%Parceiro Indicado%'` (40 hits) | WAKE \| Time Agência, 2025-12-16 ("isso não é diferenciado por peso") |
| Distribuição padrão = 3 agências por rodada | ✓ verificado | `LIKE '%3 agencias%' OR contexto Sapataria Nova` | WAKE \| Time Agência, 2025-07-11 |
| Indicações "resetam" e agências voltam ao carrossel | ✓ verificado | `LIKE '%resetaram%'` | WAKE \| Time Agência, 2025-07-04 |
| Renato (Partner Ops) é o ponto inicial para problemas Community | ✓ verificado | `LIKE '%Community%' AND LIKE '%Renato%'` | Parcerias Dotkom & Wake, 2025-07-14; WAKE \| Time Agência, 2025-07-14 |
| Lucas (time SF) responsável por correções estruturais Community | ✓ verificado | `LIKE '%Lucas%' AND LIKE '%Community%'` | WAKE \| Time Agência, 2025-07-24 |
| Procedimento de indicação manual: remover Operador + preencher Indicados 1/2/3 | ✓ verificado | `LIKE '%Parceiro Indicado 1%' AND LIKE '%retirar%'` | WAKE \| Time Agência, 2025-07-24 (caso Aste textual) |
| Reset de indicação errada exige ticket Sales Ops + remoção do dashboard de contagem | ✓ verificado | `LIKE '%Sales Ops%' OR LIKE '%dashboard%'` | WAKE \| Time Agência, 2025-07-22, 2025-07-23 |
| Conta teste de agência usada como Operador para bloquear carrossel | ✓ verificado | `LIKE '%conta Teste%' OR contexto 2025-07-03` | WAKE \| Time Agência, 2025-07-03 |
| Pacote Community = 60 licenças (~30 em uso ativo em jan/2026) | ✓ verificado | `LIKE '%licença%' AND LIKE '%Community%'` | WAKE \| Time Agência, 2026-01-16 |
| Lead via Community pode não aparecer no dashboard se Origem ≠ "Parcerias e Canais" exata | ✓ verificado | `LIKE '%Community%' AND LIKE '%Dash%'` | WAKE \| Time Agência, 2025-08-15 |
| Sugestão de mover campo Parceiro Operador para aba Parcerias (status de deploy não confirmado) | ⚠ inferido | sugestão registrada, sem confirmação de implementação | WAKE \| Time Agência, 2025-07-24 — não há mensagem posterior confirmando go-live da mudança |
| Habilitação de leads no Community condicionada a certificações mínimas | ⚠ inferido | pergunta aberta, sem decisão registrada | WAKE \| Time Agência, 2025-10-09 — questão levantada sem mensagem de resposta no corpus consultado |
| Existência de Usuário Master único por agência no Community | ⚠ inferido | resposta indica "creio que" sem confirmação técnica | Simples Inovação & Wake Parcerias, 2025-08-26 |
| SLA de resposta Sales Ops para tickets de reset de indicação | ❌ sem fonte | `LIKE '%SLA%'` em grupos relevantes sem hits específicos | Wake não publica SLA formal a parceiros (ver `wake-partnership-kb-playbook.md` §4) |

**Consultas SQLite executadas** (total: 6 queries principais):

1. `LIKE '%carrossel%'` filtrado por grupos Agência/Parcerias/Coordenação — 40 hits relevantes.
2. `LIKE '%Parceiro Operador%' OR LIKE '%Parceiro Indicado%'` — 40 hits ordenados por timestamp.
3. `LIKE '%Community%'` filtrado por grupos relevantes — 30 hits.
4. `LIKE '%Samires%' OR LIKE '%renovaç%' OR LIKE '%Venda de Canal%' OR LIKE '%Indicação Agências%' OR LIKE '%Indicação Parcerias%'` — 40 hits cobrindo bug de renovação e taxonomia.
5. `LIKE '%Sales Ops%' OR LIKE '%resetar%' OR LIKE '%aba Parcerias%' OR LIKE '%origem do lead%'` — 30 hits sobre processo de reset.
6. `LIKE '%OPT%' AND (LIKE '%qualifica%' OR LIKE '%fechamento%' OR LIKE '%converti%' OR LIKE '%novo contrato%')` — 20 hits sobre ciclo de vida da OPT.

Todas as consultas filtradas por `c.jid LIKE '%@g.us'` e `(m.media_type IS NULL OR m.media_type = '')` conforme `wake-partnership-kb-playbook.md` §7.

**Tags do artigo**: `salesforce`, `carrossel`, `lead`, `pipeline`, `agência`, `programa-parceiros`, `audience-pam`, `audience-sales`
