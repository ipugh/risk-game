console.debug("models.js file loaded");
// Territory object
function Territory(x, y, id, surrounding=[], troops=0, player="") {
    this.x = x;
    this.y = y;
    this.id = id;
    this.surrounding = surrounding;
    this.troops = troops;
    this.player = player;
}

Territory.prototype.toString = function() {
    return `id: ${this.id}, (${this.x}, ${this.y}), surrounding: ${this.surrounding}, troops: ${this.troops}, player: ${this.player}`
}

// World object
function World() {
    this.territories = {};
}

// This is where the territories are created
World.prototype.generateTerritories = function() {
    this.territories[0] = new Territory(50, 50, 0, [1]);
    this.territories[1] = new Territory(100, 100, 1, [0]);
}

// Draw the world, territories and edges
World.prototype.drawWorld = function(ctx) {
    let territorySize = 10;
    for (let [id, territory] of Object.entries(this.territories)) {
        console.debug(territory)

        // draw the territory
        ctx.beginPath();
        ctx.arc(territory.x, territory.y, territorySize, 0, 2*Math.PI);
        ctx.stroke();

        // draw surrounding
        // TODO: math so lines don't run into circles
        territory.surrounding.forEach(e => {
            ctx.beginPath();

            // draw edge from id to e
            ctx.moveTo(territory.x, territory.y);
            ex = this.territories[e].x;
            ey = this.territories[e].y;
            ctx.lineTo(ex, ey);
            ctx.stroke();
        })
    }
}
