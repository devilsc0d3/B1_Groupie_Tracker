var range = document.getElementById("range");
var rangeValue = document.getElementById("rangeValue");
rangeValue.innerHTML = range.value;
range.oninput = function() {
    rangeValue.innerHTML = this.value;
}