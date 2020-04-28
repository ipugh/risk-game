console.debug("models.js file loaded");

var territorySize = 10;

// Territory object
function Territory(x, y, id, name="unnamed", surrounding=[], troops=0, player="") {
    this.x = x;
    this.y = y;
    this.id = id;
    this.name = name;
    this.surrounding = surrounding;
    this.troops = troops;
    this.player = player;
}

Territory.prototype.toString = function() {
    return `name: ${this.name}, id: ${this.id}, (${this.x}, ${this.y}), surrounding: ${this.surrounding}, troops: ${this.troops}, player: ${this.player}`
}

// World object
function World() {
    this.territories = {};
}

// This is where the territories are created
World.prototype.generateTerritories = function() {
    this.territories[0] = new Territory(50, 50, 0, "Western United States", [1]);
    this.territories[1] = new Territory(100, 100, 1, "Eastern United States", [0]);
}

// Draw the world, territories and edges
World.prototype.drawWorld = function(ctx) {
    drawnEdges = {}

    // draw edges
    for (let [id, territory] of Object.entries(this.territories)) {
        territory.surrounding.forEach(e => {
            // Check if edge is already drawn or not
            let eHas = (drawnEdges[e] != undefined && drawnEdges[e].indexOf(id.toString()) != -1)
            let idHas = (drawnEdges[id] != undefined && drawnEdges[id].indexOf(e.toString()) != -1)
            if (!(eHas || idHas)) {
                ctx.beginPath();

                // draw edge from id to e
                ctx.moveTo(territory.x, territory.y);
                ex = this.territories[e].x;
                ey = this.territories[e].y;
                ctx.lineTo(ex, ey);
                ctx.strokeStyle = "Red"
                ctx.stroke();

                // Add edge to drawn edges
                if (drawnEdges[e] == undefined) {
                    drawnEdges[e] = []
                }
                drawnEdges[e].push(id)
            }
        });
    }

    // draw territories
    for (let [id, territory] of Object.entries(this.territories)) {
        console.debug(territory)

        // draw the territory
        ctx.beginPath();
        ctx.arc(territory.x, territory.y, territorySize, 0, 2*Math.PI);
        ctx.fillStyle="Black"
        ctx.fill();
    }
}

// return id onclick, -1 if none clicked
World.prototype.territoryIsClicked = function(x, y) {
    let dist = (xa, ya, xb, yb) => Math.sqrt(((xa-xb)**2) + ((ya-yb)**2))
    for (let [id, territory] of Object.entries(this.territories)) {
        if (dist(x, y, territory.x, territory.y) <= territorySize) {
            return id
        }
    }
    return -1
}
