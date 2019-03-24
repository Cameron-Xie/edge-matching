module.exports = class Puzzle {
    constructor(left, top, right, bottom) {
        this.edges = [left, top, right, bottom];
        this.origin = [left, top, right, bottom];
        this.active = true;
        this.rotation = 0;
        this.atFirstPosition = 0;
    }

    get left() {
        return this.edges[0];
    }

    get top() {
        return this.edges[1];
    }

    get right() {
        return this.edges[2];
    }

    get bottom() {
        return this.edges[3];
    }

    Rotate() {
        const firstEdge = this.edges.shift();
        this.edges.push(firstEdge);
        this.rotation++;
    }

    Reset() {
        this.edges = [...this.origin];
        this.active = true;
        this.rotation = 0;
    }
}