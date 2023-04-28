const fs = require('fs');
const path = require('path');
const simpleGit = require('simple-git');
const matter = require('gray-matter');
const { Console } = require('console');

const git = simpleGit();
const green = '\x1b[32m';
const reset = '\x1b[0m';
const yellow = '\x1b[33m';


const MASTER_RECORD_PATH = path.join(__dirname, '../docs/README.md');

/**
 * TODO:
 *  - Need to update master record when a section that had front matter is deleted.
 *  - Handle renames, deleted old record and a new section.
 */
async function main() {
    try {
        // Get the list of staged files.
        const stagedFiles = await git.diff(['--name-only', '--cached']);

        // Filter the staged files for .md files.
        const mdFiles = stagedFiles
            .split('\n')
            .filter(file => file.endsWith('.md'));

        // Read the master record and extract just the front matter.
        let { data, content } = matter(fs.readFileSync(MASTER_RECORD_PATH, 'utf8'));

        // Create a deep clone of the master record for comparison later
        const clonedMasterRecord = structuredClone(data);

        let masterRecordChanged = false;

        for (const mdFile of mdFiles) {
            // Read the content of the .md file.
            const fileContent = fs.readFileSync(mdFile, 'utf8');

            // Parse the YAML front matter.
            const parsedExercise = matter(fileContent);

            const exerciseData = parsedExercise.data

            // merge the front-matter from the staged file into the master record.
            // This will overwrite the inner object on the master record if it exists
            // and also add it if it doesnt.
            if (exerciseData) {
                data = { ...data, ...exerciseData };
                masterRecordChanged = true;
            }
        }

        // Create an array of key-value pairs and then sorts them alphabetically
        // by the keys.
        const sortedEntries = Object.entries(data).sort(
            (a, b) => a[0].localeCompare(b[0])
        );
        const sortedMasterEntries = Object.entries(clonedMasterRecord).sort(
            (a, b) => a[0].localeCompare(b[0])
        );

        // Create a new object from the sorted entries
        const sortedData = Object.fromEntries(sortedEntries);
        const sortedMasterRecord = Object.fromEntries(sortedMasterEntries);

        // Stringify the object and compare. This handles the case of an .md file being updated
        // but not the front-data
        if (JSON.stringify(sortedData) === JSON.stringify(sortedMasterRecord)) {
            masterRecordChanged = false;
        }

        if (masterRecordChanged) {
            // Stringify the updated front-matter and origional content
            const updatedContent = matter.stringify(content, sortedData);

            // Update the master record with the new front-matter. Use the sorted
            fs.writeFileSync(MASTER_RECORD_PATH, updatedContent);

            console.error(`${yellow}New front matter detected, master record updated${reset}`);
            console.error(`${yellow}Please review changes to ./docs/README.md${reset}`);
            process.exit(1);
        } else {
            // Continue with the commit.
            console.log(`\n\n${green}No changes to master record, proceeding with commit.${reset}`);
            process.exit(0);
        }

    } catch (error) {
        console.error('An error occurred:', error);
        process.exit(1);
    }
}

main();