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
    let startPosition;
    for (let i = 0; i < nRows; i++) {
        for (let j = 0; j < nColumns; j++) {
            puzzle[i][j].position = [i, j];

            if (puzzle[i][j].current) {
                startPosition = puzzle[i][j];
            }
        }
    }

    const copyPuzzle = (puzzle) => {
        const copy = [];
        for (let i = 0; i < puzzle.length; i++) {
            const row = [];
            for (let j = 0; j < puzzle[0].length; j++) {
                row.push({ ...puzzle[i][j] });
            }
            copy.push(row);
        }
        return copy;
    };

    const runSimulation = (puzzle, currentPosition, maxSteps = 10000) => {
        let outside = false;
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

        // check visited
        let visited = [];
        for (let i = 0; i < nRows; i++) {
            for (let j = 0; j < nColumns; j++) {
                if (puzzle[i][j].visited) {
                    visited.push([i, j]);
                }
            }
        }

        return [visited, n + 1];
    };

    //

    let currentPosition = { ...startPosition };
    let puzzleRun = copyPuzzle(puzzle);

    let [visited, n] = runSimulation(puzzleRun, currentPosition);
    console.log("Number of steps", n);
    console.log("Visited", visited.length);

    //

    const maxSteps = 100000;

    //console.log(visited);

    const obstructions = [];
    let currentStep = 1;
    const maxStep = visited.length;
    const logEachStep = 100;
    for (visitedPosition of visited) {
        let puzzleRun = copyPuzzle(puzzle);
        let currentPosition = { ...startPosition };

        puzzleRun[visitedPosition[0]][visitedPosition[1]].obstacle = true;

        let [_, n] = runSimulation(puzzleRun, currentPosition, maxSteps);

        if (n >= maxSteps) {
            obstructions.push(visitedPosition);
        }

        if (currentStep % logEachStep === 0) {
            console.log(`Step ${currentStep}/${maxStep}`);
        }

        currentStep++;
    }

    console.log("Total obstructions", obstructions.length);
});
