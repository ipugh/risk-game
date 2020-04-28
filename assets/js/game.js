console.debug("game.js file loaded");

var w = new World();
w.generateTerritories();


// document on ready
$(function() {
    var c = document.getElementById("gameCanvas");
    leftoffset = c.offsetLeft + c.clientLeft;
    topoffset = c.offsetTop + c.clientTop;

    // print to console on click
    c.addEventListener('click', function(event) {
        let x = event.pageX - leftoffset;
        let y = event.pageY - topoffset;
        console.log(`click event at (x: ${x}, y: ${y})`);

        let territory = w.territoryIsClicked(x, y)
        console.log(`territory clicked: ${territory}`)

    });

    // you can draw stuff to canvas with ctx
    var ctx = c.getContext("2d");
    w.drawWorld(ctx)
});


