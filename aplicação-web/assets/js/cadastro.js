$('#formulario-cadastro').on('submit', criarUsuario); //quando um formulario com esse nome receber um submit ele faz essa função
// $ -> pega dados

function criarUsuario(evento){
    evento.preventDefault();

    if($('#senha').val() != $('#confirmar-senha').val()){ //para pegar valor do elemento $('referencia ao elemento').val
        alert("As senhas não coincidem!");
        return
    }

    $.ajax({
        url: "/usuarios",
        method: "POST",
        data: {
            nome: $('#nome').val(),
            sobrenome: $('#sobrenome').val(), // Se necessário
            login: $('#login').val(),
            nick: $('#nick').val(),
            senha: $('#senha').val(),
            perm: $('#perm').val() // Adicionando permissão
        }
    }).done(function() {
        alert("Usuário cadastrado com sucesso!");
    }).fail(function(erro) {
        alert(erro);
    });
    
}