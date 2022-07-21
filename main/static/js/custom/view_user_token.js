
$(function () {
    console.log("vm", vm)
    document.getElementById("UserToken").className += " active";
    var mainArray = [];
    var table = "";

    function createActiveArray(values, keys) {
        var subArray = [];
        for (i = 0; i < values.length; i++) {
            for (var propertyName in values[i]) {
                subArray.push(values[i][propertyName]);
            }
            subArray.push(keys[i])
            mainArray.push(subArray);
            subArray = [];
        }
    }

    /*Function for assigning data array into data table*/
    function dataTableManipulate() {
        table = $("#user_token_table").DataTable({
            data: mainArray,
            "ordering": false,
            "pageLength": 50,
            "columnDefs": [
                {
                    "targets": -1,
                    "width": "10%",
                    "data": null,
                    "defaultContent": '<div class="edit-wrapper"><span class="icn"><i class="fa fa-clipboard" aria-hidden="true" id="tokenCopy" title="Token Copy" style = "color:red;cursor: pointer;"></i><i class="fa fa-trash-o" title="Recall Token" aria-hidden="true" id="delete" style = "color:red;cursor: pointer;margin-left: 30px;"></i></span></div>'


                }]
        });
        /*Add a plus symbol in webpage for add new groups*/
        var item = $('<span style="padding-left:37px"><input type="image"  style ="height: 28px;    width: 46px;margin-left: 150px;" src="/static/img/add2.png" ></span>');
        item.click(function () {
            window.location = '/userToken/add';
        });
        $('.table-responsive .dataTables_filter').append(item);
    }

    if (vm.Values != null) {
        createActiveArray(vm.Values, vm.Keys);
    }
    dataTableManipulate();


    $('#user_token_table tbody').on('click', '#delete', function () {
        var data = table.row($(this).parents('tr')).data();
        var userTokenID = data[5];
        $("#myModal").modal();
        $("#recallOk").click(function () {
            $.ajax({
                type: "POST",
                url: '/userToken/recall/' + userTokenID,
                data: '',
                success: function (feedback) {
                    if (feedback == "true") {
                        window.location = '/userToken'
                    }
                }
            });
        });

    });
    $('#user_token_table tbody').on('click', '#tokenCopy', function () {
        var data = table.row($(this).parents('tr')).data();
        var userToken = data[2];
        console.log("userToken", userToken);
        document.getElementById('userTokenData').value = userToken;
        $("#tokenCopyModal").modal();

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
