$(function () {
    var username = "superAdmin"
    $("#signIn").click(function () {
        $.ajax({
            type: 'POST',
            url: '/',
            data: {
                'username': $("#username").val(),
                'password': $("#password").val()
            },
            success: function (data) {
                var jsonData = JSON.parse(data);
                console.log(jsonData[0]);
                if (jsonData[0] == "true") {
                    window.location = "/dashboard";
                } else {
                    $("#login_err").css({ "color": "red", "font-size": "15px", "margin-left": "67px" });
                    $("#login_err").html("Invalid Username or Password!");
                }
            }
        });

        return false;
    });

});
