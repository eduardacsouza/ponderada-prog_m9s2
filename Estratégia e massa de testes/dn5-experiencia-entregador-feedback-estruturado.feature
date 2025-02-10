#language: pt
Funcionalidade: Coleta e análise de feedback dos entregadores
  Eu como administrador do sistema
  Quero coletar feedback dos entregadores
  Para melhorar a experiência de uso do aplicativo

  Cenário: Coleta de feedback bem-sucedida
    Dado que um entregador acessa o formulário de feedback dentro do app
    Quando ele preenche e envia a resposta
    Então a resposta deve ser armazenada corretamente no sistema

  Cenário: Monitoramento da taxa de resposta
    Dado que o sistema analisa a taxa de resposta
    Quando ela cair abaixo de 30%
    Então um alerta deve ser gerado automaticamente

  Cenário: Análise de tendências de churn e insatisfação
    Dado que os feedbacks são coletados mensalmente
    Quando o sistema processa os dados
    Então ele deve identificar padrões de churn e insatisfação
