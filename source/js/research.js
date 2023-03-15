
var firstA = document.getElementById("firstA");
var firstAValue = document.getElementById("firstAValue");
firstAValue.innerHTML = firstA.value;
firstA.oninput = function() {
    firstAValue.innerHTML = this.value;
}

var creationD = document.getElementById("creationD");
var creationDValue = document.getElementById("creationDValue");
creationDValue.innerHTML = creationD.value;
creationD.oninput = function() {
    creationDValue.innerHTML = this.value;
}