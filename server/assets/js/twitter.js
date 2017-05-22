$.getJSON("/tweets", function(data) {
    var items = [];
    $.each(data, function(key, val) {
        items.push(
		"<div class=tweet id='tweet" + key + "'>" + val +
		"<br>" +
		"<button class='btn btn-primary'>tweet</button>" +
		"<button class='btn btn-danger'>blacklist</button>" +
		"</div>"
	);
    });

    $("<div/>", {
        "class": "app",
        "contenteditable": "true",
        html: items.join("")
    }).appendTo("body");
});
