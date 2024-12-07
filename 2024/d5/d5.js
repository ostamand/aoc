const fs = require("node:fs");
const os = require("os");

fs.readFile("input.txt", { encoding: "utf8" }, (err, data) => {
    // read input file
    const lines = data.split(os.EOL);

    const orderingRules = [];
    const updates = [];
    let readingRules = true;
    for (let i = 0; i < lines.length; i++) {
        const line = lines[i];
        if (line === "") {
            readingRules = false;
        } else {
            if (readingRules) {
                orderingRules.push(line.split("|").map((rule) => Number(rule)));
            } else {
                updates.push(line.split(",").map((update) => Number(update)));
            }
        }
    }

    // create set with all numbers
    const numbers = new Set();
    for (rule of orderingRules) {
        numbers.add(rule[0]);
        numbers.add(rule[1]);
    }

    // extract rules by number
    const rulesByNumber = new Map();
    for (number of numbers) {
        const rules = {
            before: [],
            after: [],
        };
        for (rule of orderingRules) {
            if (rule[0] === number) {
                rules.after.push(rule[1]);
            } else if (rule[1] === number) {
                rules.before.push(rule[0]);
            }
        }
        rulesByNumber.set(number, rules);
    }

    // check which update is valid
    const validUpdates = [];
    for (update of updates) {
        let isValid = true;

        for (let i = 0; i < update.length; i++) {
            const before = update.filter((_, index) => index < i);
            const after = update.filter((_, index) => index > i);

            const currentNumber = update[i];

            for (numberBefore of before) {
                if (
                    rulesByNumber
                        .get(currentNumber)
                        ?.after.includes(numberBefore)
                ) {
                    isValid = false;
                }
            }
            for (numberAfter of after) {
                if (
                    rulesByNumber
                        .get(currentNumber)
                        ?.before.includes(numberAfter)
                ) {
                    isValid = false;
                }
            }
        }
        if (isValid) {
            validUpdates.push(update);
        }
    }

    // find middle page numbers
    let total = 0;
    for (validUpdate of validUpdates) {
        const middleIndex = Math.floor(validUpdate.length / 2);
        total += validUpdate[middleIndex];
    }

    console.log("total", total);
});
