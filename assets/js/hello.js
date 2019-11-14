$(document).ready(function () {
    $.ajax({
        url: "/api/test",
        success: function (data) {
            data = JSON.parse(data)
            $('.greeting-id').append(data.id);
            $('.greeting-content').append(data.content);
        }
    })
});