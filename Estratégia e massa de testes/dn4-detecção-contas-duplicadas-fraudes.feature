#language: pt
Funcionalidade: Detecção de contas duplicadas e fraudes
  Eu como administrador do sistema
  Quero impedir a criação de contas fraudulentas
  Para garantir a segurança da plataforma

  Cenário: Cadastro com CPF banido
    Dado que um usuário tenta criar uma conta com um CPF banido
    Quando a verificação for realizada
    Então o sistema deve rejeitar o cadastro e exibir um aviso ao usuário

  Cenário: Cadastro com CPF válido
    Dado que um usuário cria uma conta com um CPF nunca utilizado
    Quando a verificação for realizada
    Então a conta deve ser criada com sucesso

  Cenário: Tempo de resposta do sistema
    Dado um cadastro legítimo
    Quando o sistema processa a validação
    Então o tempo de resposta deve ser inferior a 3 segundos
