# book-api

1. cadastrar um usuário (createUser)
2. cadastrar livro para um usuário (addBookToMyCollection)
3. emprestar meu livro para outro usuário (lendBook)
4. devolver o livro (returnBook)
5. pegar detalhes de um usuário (user)

Regras de negócio

● Em alguns lugares você vai encontrar o input loggedUserId isso representa o ID do
usuário logado, por simplficação do case, para não ter que implementar autenticação,
guardar senhas, gerar e parsear tokens nas requisições, optamos por deixar
parametrizável. Dito isso, todas as validações e verificações no serviço que julgar
necessária devem ser feitas em relação ao usuário logado, faz parte da avaliação,
nem todas estão explicitamentes descritas aqui.
● 1. cadastrar um usuário (createUser)
○ email deve ser único no sistema
● 3. emprestar meu livro para outro usuário (lendBook)
○ Apenas um empréstimo por livro de um usuário por vez deve ser permitido, ou
seja, enquanto não for devolvido o livro emprestado este não poderá ser
emprestado

● 4. devolver o livro (returnBook)
○ O livro só pode ser devolvido uma vez. Ao tentar devolver um livro já devolvido,
um erro deve ser retornado
● 5. pegar detalhes de um usuário (user)
○ Não há necessidade de guardar histórico de empréstimos. Mas depedendo da
estrutura de dados escolhida, já pode estar implementada. Ou seja, por
simplificação, você pode retornar apenas 1 empréstimo por livro ou o histórico
inteiro. Apenas não considerar o histórico de empréstimos uma obrigatoriedade.