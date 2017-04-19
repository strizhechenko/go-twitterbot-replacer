$.getJSON("http://127.0.0.1:8080/tweets", function(data) {
    var items = [];
    $.each(data, function(key, val) {
        items.push("<div class=tweet id='tweet" + key + "'>" + val + "</div>");
    });

    $("<div/>", {
        "class": "app",
        "contenteditable": "true",
        html: items.join("")
    }).appendTo("body");
});
