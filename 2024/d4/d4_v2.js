const fs = require("node:fs");
const os = require("os");

fs.readFile("input.txt", "utf-8", (error, data) => {
    const words = new Set();

    const wordSearch = data.split(os.EOL).map((line) => line.split(""));

    const nRows = wordSearch.length;
    const nColumns = wordSearch[0].length;

    // add index for reference
    for (let i = 0; i < nRows; i++) {
        for (let j = 0; j < nColumns; j++) {
            wordSearch[i][j] = {
                index: `[${i},${j}]`,
                value: wordSearch[i][j],
            };
        }
    }

    const invertLine = (line) => {
        const inverted = [];
        for (let i = line.length; i >= 0; i--) {
            inverted.push(line[i]);
        }
        return inverted;
    };

    const processLine = (line) => {
        let i = 0;
        const n = line.length;
        while (i < n) {
            if (
                line[i]?.value === "X" &&
                line[i + 1]?.value === "M" &&
                line[i + 2]?.value === "A" &&
                line[i + 3]?.value === "S"
            ) {
                const index = `${line[i].index}-${line[i + 3].index}`;
                // might not be needed
                const indexInverted = `${line[i + 3].index}-${line[i].index}`;
                if (!words.has(index) && !words.has(indexInverted)) {
                    words.add(index);
                }
                i += 4;
            } else {
                i++;
            }
        }
    };

    // horizontal rows
    for (let i = 0; i < nRows; i++) {
        processLine(wordSearch[i]);
        processLine(invertLine(wordSearch[i]));
    }

    // vertical rows
    for (let j = 0; j < nColumns; j++) {
        const column = [];
        for (let i = 0; i < nRows; i++) {
            column.push(wordSearch[i][j]);
        }
        processLine(column);
        processLine(invertLine(column));
    }

    // diagonals
    for (let i = 0; i < nRows; i++) {
        for (let j = 0; j < nColumns; j++) {
            let diagonal = [];
            let n = i;
            let m = j;

            // diagonal down
            while (n >= 0 && m >= 0 && n < nRows && m < nColumns) {
                diagonal.push(wordSearch[n][m]);
                n++;
                m++;
            }
            processLine(diagonal);
            processLine(invertLine(diagonal));

            // diagonal up
            diagonal = [];
            n = i;
            m = j;
            while (n >= 0 && m >= 0 && n < nRows && m < nColumns) {
                diagonal.push(wordSearch[n][m]);
                n++;
                m--;
            }
            processLine(diagonal);
            processLine(invertLine(diagonal));
        }
    }

    console.log(words.size);
});
