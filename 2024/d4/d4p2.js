const fs = require("node:fs");
const os = require("os");

fs.readFile("input.txt", "utf8", (error, data) => {
    const wordSearch = data.split(os.EOL).map((line) => line.split(""));

    const nRows = wordSearch.length;
    const nColumns = wordSearch[0].length;

    let totalCount = 0;
    for (let i = 0; i < nRows; i++) {
        for (let j = 0; j < nColumns; j++) {
            if (wordSearch[i][j] === "A") {
                // one diagonal
                if (
                    i - 1 >= 0 &&
                    i + 1 < nRows &&
                    j - 1 >= 0 &&
                    j + 1 < nColumns &&
                    ((wordSearch[i - 1][j - 1] === "M" &&
                        wordSearch[i + 1][j + 1] === "S") ||
                        (wordSearch[i - 1][j - 1] === "S" &&
                            wordSearch[i + 1][j + 1] === "M"))
                ) {
                    if (
                        i - 1 >= 0 &&
                        i + 1 < nRows &&
                        j - 1 >= 0 &&
                        j + 1 < nColumns &&
                        ((wordSearch[i - 1][j + 1] === "M" &&
                            wordSearch[i + 1][j - 1] === "S") ||
                            (wordSearch[i - 1][j + 1] === "S" &&
                                wordSearch[i + 1][j - 1] === "M"))
                    ) {
                        totalCount += 1;
                    }
                }
            }
        }
    }

    console.log(totalCount);
});
