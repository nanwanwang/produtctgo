$(document).ready(function(){
    $("#login-form").submit(function (event){
        event.preventDefault();
        $.ajax({
            type:'POST',
            url:'/login',
            data: {
                'username': $('input[name=username]').val(),
                'password': $('input[name=password]').val()
            },
            dataType: 'json',
            encode: true
        }).done(function (data) {
            console.log(data);
            if (data.code === 1) {
                window.location.href = '/';
            } else {
                alert(data.message)
            }

        }).fail(function (err) {
            console.log(err)
            alert(data.statusText)
        })
    });
})