# 1. Mapa Business Drivers

## DN4 - Detecção de Contas Duplicadas e Fraudes
**Requisito**: O sistema deve identificar e bloquear 95% das tentativas de criação de contas utilizando CPFs já associados a contas banidas.

**Métrica**: Tempo de resposta para validar documentos e históricos de banimentos deve ser inferior a 3 segundos.

**Monitoramento**: Se a taxa de aprovação de contas suspeitas ultrapassar 5%, um alerta deve ser disparado para revisão manual.

## DN5 - Experiência do Entregador e Feedback Estruturado
**Requisito**: O sistema deve oferecer um mecanismo nativo de feedback dentro do app, garantindo que pelo menos 50% dos entregadores ativos forneçam feedback mensalmente.

**Métrica**: Respostas aos formulários devem ser coletadas automaticamente e analisadas com base em tendências de churn e insatisfação.

**Monitoramento**: Se a taxa de resposta cair abaixo de 30%, um alerta deve ser gerado para realização de medidas que incentivem a participação das pesquisas.

# 2. Estratégia e massa de testes
&emsp; Os arquivos correspondentes as estratégias e massas de testes se encontram em: 
[Estratégia e Massa de Testes](https://github.com/eduardacsouza/ponderada-prog_m9s2/tree/main/Estratégia%20e%20massa%20de%20testes)

# 3.Codificação como documentação de testes
&emsp; O arquivo correspondente a codificação como documentação de testes se encontra em:
[Codificação como documentação de testes](https://github.com/eduardacsouza/ponderada-prog_m9s2/blob/main/main_test.go)






