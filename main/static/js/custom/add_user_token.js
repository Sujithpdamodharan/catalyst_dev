document.getElementById("UserToken").className += " active";

$().ready(function () {

    var repeat = "";

    $("#addUserTokenFormId").validate({
        rules: {
            addUserName: "required",
            addMobile: "required"
        },
        messages: {

        },
        submitHandler: function () {
            $("#addButton").attr('disabled', true);
            var formData = $("#addUserTokenFormId").serialize();
            $.ajax({
                url: '/userToken/add',
                type: 'post',
                datatype: 'json',
                data: formData,
                //call back or get response here
                success: function (response) {
                    var jsonData = JSON.parse(response);
                    console.log("jsonData", jsonData[0])
                    if (jsonData[0] == "true") {
                        document.getElementById('userTokenData').value = jsonData[1];
                        $('#tokenCopyModal').modal({
                            backdrop: 'static',
                            keyboard: true,
                            show: true
                        });
                        $("#tokenCopyModal").modal();

                        $('#closemodal').click(function () {
                            window.location = "/userToken"
                        });

                    } else {
                        window.location = "/";
                    }
                },
                error: function (request, status, error) {
                }
            });
        }
    });
});


function copyToken() {
    /* Get the text field */
    var copyText = document.getElementById("userTokenData");

    /* Select the text field */
    copyText.select();
    copyText.setSelectionRange(0, 99999); /* For mobile devices */

    /* Copy the text inside the text field */
    navigator.clipboard.writeText(copyText.value);
}

