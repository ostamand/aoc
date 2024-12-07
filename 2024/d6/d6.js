const fs = require("node:fs");
const os = require("os");

fs.readFile("input.txt", { encoding: "utf8" }, (err, data) => {
    const puzzle = data.split(os.EOL).map((row) =>
        row.split("").map((cell) => {
            const cellData = {
                visited: false,
                current: false,
                obstacle: false,
            };
            if (cell === "^" || cell === ">" || cell === "<" || cell === "v") {
                cellData.visited = true;
                cellData.current = true;
                if (cell === "^") {
                    cellData.direction = [-1, 0];
                }
                if (cell === ">") {
                    cellData.direction = [0, 1];
                }
                if (cell === "<") {
                    cellData.direction = [0, -1];
                }
                if (cell === "v") {
                    cellData.direction = [1, 0];
                }
            }
            if (cell === "#") {
                cellData.obstacle = true;
            }
            return cellData;
        })
    );

    const nRows = puzzle.length;
    const nColumns = puzzle[0].length;

    // fill in index
    let currentPosition;
    for (let i = 0; i < nRows; i++) {
        for (let j = 0; j < nColumns; j++) {
            puzzle[i][j].position = [i, j];

            if (puzzle[i][j].current) {
                currentPosition = puzzle[i][j];
            }
        }
    }

    let outside = false;
    let maxSteps = 10000;
    let n = 0;
    while (!outside && n < maxSteps) {
        const nextPositionIndex = [
            currentPosition.position[0] + currentPosition.direction[0],
            currentPosition.position[1] + currentPosition.direction[1],
        ];
        // check if obstacle @ next position
        if (
            nextPositionIndex[0] < 0 ||
            nextPositionIndex[1] < 0 ||
            nextPositionIndex[0] >= nRows ||
            nextPositionIndex[1] >= nColumns
        ) {
            outside = true;
        } else {
            const nextPosition =
                puzzle[nextPositionIndex[0]][nextPositionIndex[1]];
            if (!nextPosition.obstacle) {
                nextPosition.visited = true;
                nextPosition.direction = currentPosition.direction;
                currentPosition = nextPosition;
            } else {
                // need to turn to the right
                if (
                    currentPosition.direction[0] == -1 &&
                    currentPosition.direction[1] == 0
                ) {
                    // "^"
                    currentPosition.direction = [0, 1]; // >
                } else if (
                    currentPosition.direction[0] == 0 &&
                    currentPosition.direction[1] == 1
                ) {
                    // >
                    currentPosition.direction = [1, 0]; // v
                } else if (
                    currentPosition.direction[0] == 1 &&
                    currentPosition.direction[1] == 0
                ) {
                    // v
                    currentPosition.direction = [0, -1]; // <
                } else if (
                    currentPosition.direction[0] == 0 &&
                    currentPosition.direction[1] == -1
                ) {
                    // <
                    currentPosition.direction = [-1, 0]; // ^
                }
            }
        }
        n++;
    }

    console.log("Number of steps", n);

    let totalVisited = 0;
    for (let i = 0; i < nRows; i++) {
        for (let j = 0; j < nColumns; j++) {
            if (puzzle[i][j].visited) {
                totalVisited += 1;
            }
        }
    }

    console.log("Visited", totalVisited);
});
