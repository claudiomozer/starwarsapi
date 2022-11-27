# Carregar um planeta da API através do id

> ## Casos de sucesso

1. ✅ Recebe uma requisição **POST** no endpoint **/api/planets**
2. ✅ Valida se foi recebido o parametro **id** no body da requisição
3. ✅ Verifica se o planeta com determinada URL (URL única da API com id) já não foi importando
4. ✅ Busca o planeta por id na API externa <
5. ✅ Verifica se os filmes com as URLs vinculadas ao planeta
6. ✅ **Cria** o filme caso não existir na base de dados
7. ✅ **Cria** o planeta no banco de dados
8. ✅ Retorna 201 com o **id** da base de dados

> ## Erros

1. ❌ Retorna 400 caso o **id** não seja fornecido
2. ❌ Retorna 302 caso o planeta já exista na base de dados
3. ❌ Retorna 404 caso o planeta não tenha sido encontrado na API externa
4. ❌ Retorna 500 em casos de timeouts com a API externa
