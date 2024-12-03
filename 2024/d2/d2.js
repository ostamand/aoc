const fs = require("node:fs");
var os = require("os");

const processLevels = (levels) => {
  let countSafeLevel = 0;
  for (level of levels) {
    const levelDifferences = level
      .map((value, index) => level[index + 1] - value)
      .filter((value) => !isNaN(value));

    // all increasing
    const allIncreasing = levelDifferences.every((value) => value > 0);
    const allDecreasing = levelDifferences.every((value) => value < 0);
    const allDifferences = levelDifferences
      .map((value) => Math.abs(value))
      .every((value) => value >= 1 && value <= 3);

    countSafeLevel += (allIncreasing || allDecreasing) && allDifferences;
  }

  console.log("Number of safe levels:", countSafeLevel);
};

const exampleData = [
  [7, 6, 4, 2, 1],
  [1, 2, 7, 8, 9],
  [9, 7, 6, 2, 1],
  [1, 3, 2, 4, 5],
  [8, 6, 4, 4, 1],
  [1, 3, 6, 7, 9],
];

//processLevels(exampleData);

let levels = null;
fs.readFile("data.txt", "utf8", (err, data) => {
  if (err) {
    console.error(err);
    return;
  }
  levels = data
    .split(os.EOL)
    .map((row) => row.split(" ").map((item) => Number(item)));

  processLevels(levels);
});
