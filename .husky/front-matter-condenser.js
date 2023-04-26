const fs = require('fs');
const path = require('path');
const simpleGit = require('simple-git');
const matter = require('gray-matter');
const YAML = require('json-to-pretty-yaml');
const { Console } = require('console');

const git = simpleGit();

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
            // const exerciseData = matter(fileContent).data;
            const parsedExercise = matter(fileContent);

            // Check if we deleted a section.

            const exerciseData = parsedExercise.data

            console.log(exerciseData);
            // Need to handle the delete case. When an exercise section is deleted we should
            // Remove the whole section from the master record
            // When

            // merge the front-matter from the staged file into the master record.
            // This will overwrite the inner object on the master record if it exists
            // and also add it if it doesnt.
            if (exerciseData) {
                data = { ...data, ...exerciseData };
                masterRecordChanged = true;
            }
            console.log(data);

            //   // START HERE I want to modify this logic so that if the key of data is included in masterRecord then update
            //   // masterRecord. If not then push all of data into masterRecord
            //   // If the file has front matter and it's not in the master record, add it.
            //   if (data && !masterRecord.some(entry => JSON.stringify(entry) === JSON.stringify(data))) {
            //     masterRecord.push(data);
            //     masterRecordChanged = true;
            //   }
        }

        // chatgpt code to sort object by key
        const sortedEntries = Object.entries(data).sort(
            (a, b) => a[0].localeCompare(b[0])
        );
        const sortedMasterEntries = Object.entries(clonedMasterRecord).sort(
            (a, b) => a[0].localeCompare(b[0])
        );

        // Create a new object from the sorted entries
        const sortedData = Object.fromEntries(sortedEntries);
        const sortedMasterRecord = Object.fromEntries(sortedMasterEntries);
        // end of chatgpt code

        if (JSON.stringify(sortedData) === JSON.stringify(sortedMasterRecord)) {
            masterRecordChanged = false;
        }

        //console.log(Object.keys(masterRecord));
        //console.log(YAML.stringify(JSON.parse(fs.readFileSync('/Users/jburns/git/devops-bootcamp/docs/1-introduction/1.1-devops-defined.json', 'utf-8'))))
        const updatedContent = matter.stringify(content, data);
        // console.log(updatedContent)


        if (masterRecordChanged) {
            // We can always write the content back to the master record because git does not care about timestamps
            // So we have three cases, no .md files were staged, and .md
            fs.writeFileSync(MASTER_RECORD_PATH, updatedContent);

            console.error('New front matter detected, master record updated. Please review changes to ./docs/README.md');
            process.exit(1);
        } else {
            // Continue with the commit.
            console.log('No changes to master record, proceeding with commit.');
            process.exit(0);
        }
    } catch (error) {
        console.error('An error occurred:', error);
        process.exit(1);
    }
}

main();