$(document).ready(function(){
    $('#login').on('submit', fazerlogin);

    function fazerlogin(event) {
        event.preventDefault();
        var form = $(this);
        var dados = new FormData(this);
        $.ajax({
            url: "/login",
            type: "POST",
            data: {
                login: $('#ID1').val(),
                senha: $('#Password1').val()
            }
        }).done(function () {
            window.location.href = "/home";
        }).fail(function () {
            alert("Usuário ou senha inválidos");
        });
    }
});
