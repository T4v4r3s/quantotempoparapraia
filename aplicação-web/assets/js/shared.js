function showPassword(pass,image){
    var senha = document.getElementById(pass);
    var img = document.getElementById(image);
    if(senha.type === "password"){
      senha.type = "text";
      img.src = "/src/eye-slash-fill.svg"
    }
    else{
      senha.type = "password";
      img.src = "/src/eye-fill.svg"
    }
}