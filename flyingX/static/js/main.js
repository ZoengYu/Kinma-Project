var click_count = 1;
function onButtonClick() {
	if (click_count % 2 == 0) {
		document.getElementById('search').className = "hide";
	}
	else {
		document.getElementById('search').className = "show";
	}
	click_count = click_count + 1;
};

function loadSearch() {
	var xhttp = new XMLHttpRequest();
	xhttp.onreadystatechange = function () {
		if (this.readyState == 4 && this.status == 200) {
			displayData(this);
		}
	};
	xhttp.open("GET", "https://jsonplaceholder.typicode.com/posts/", true);
	xhttp.send();
};

function displayData(xhttp) {
	jsonData = JSON.parse(xhttp.responseText);
	var newContent = ""
	for (index in jsonData) {
		newContent += "<p>" + jsonData[index].body + "</p>";
	}
	document.getElementById("json-content").innerHTML = newContent;
};

function clearJSON() {
	document.getElementById("json-content").innerHTML = "";
}


$(function () {
	$('.carousel').carousel({
		interval: 3000
	});
});