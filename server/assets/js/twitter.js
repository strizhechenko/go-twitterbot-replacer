$.getJSON("/tweets", function(data) {
    var items = [];
    $.each(data, function(key, val) {
        items.push(
		"<div class=tweet id='tweet" + key + "'>" + val +
		"<br>" +
		"<button class='tweetbutton'>tweet</button>" +
		"<button class='blacklist'>blacklist</button>" +
		"</div>"
	);
    });

    $("<div/>", {
        "class": "app",
        "contenteditable": "true",
        html: items.join("")
    }).appendTo("body");
});
