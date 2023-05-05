/**
 * The main module for generating charts and graphs related to the bootcamp metadata.
 * @module bootcampCharts
 */

import Chart from 'chart.js/auto';
import { WordCloudController, WordElement } from 'chartjs-chart-wordcloud';
import { bootcampMetadata } from './read-metadata';

// Register the wordCloud controller, element, and scale with Chart.js.
// Starting in Chart.js 3 iirc they support 'tree shaking' which means
// componenets have to be registered explicitly so bundlers like webpack
// can effectively remove unused controllers when minifying
Chart.register(WordCloudController, WordElement);

/**
 * Generates a doughnut chart displaying the breakdown of time spent on each category of exercises in the bootcamp.
 *
 * @param {string} canvasId - The ID of the canvas element to use for the chart.
 */
function generateCategoryDoughnutChart(canvasId) {
    var canvas = document.getElementById(canvasId);
    canvas.width = 400;
    canvas.height = 400;
    const ctx = canvas.getContext('2d');

    let categoryData = {};

    // Aggrigate all time spent in each category.
    for (const doc of Object.values(bootcampMetadata)) {
        if ('category' in doc && 'estReadingMinutes' in doc) {
            // Create bucket for categories and add reading time
            if ('category' in categoryData) {
                categoryData[doc.category] += doc.estReadingMinutes;
            } else {
                categoryData[doc.category] = doc.estReadingMinutes;
            }

            if ('exercises' in doc) {
                doc.exercises.forEach(exercise => {
                    if ('estMinutes' in exercise) {
                        categoryData[doc.category] += exercise.estMinutes;
                    }
                });
            }
        }
    }

    const totalTime = Object.values(categoryData).reduce((acc, t) => { return acc + t }, 0);

    const data = {
        labels: Object.keys(categoryData),
        datasets: [{
            label: '',
            data: Object.values(categoryData).map((t) => { return t/totalTime })
        }]
    }

    const options = {
        plugins: {
            title: {
                display: true,
                text: 'Bootcamp Breakdown by Category',
                font: {
                    size: 30
                }
            },
            tooltip: {
                callbacks: {
                    label: (toolTip) => {
                        let label = toolTip.dataset.label || '';
                        let value = toolTip.formattedValue * 100 + '%';
                        return `${label}: ${value}`
                    }
                }
            }
        }
    }

    const myChart = new Chart(ctx, {
        type: 'doughnut',
        data: data,
        options: options
    });
    console.log(categoryData);
    console.log(data);
}

// TODO: Fix this method. Right now the word cloud is too small and
// I have seen it 'blow up' the page (grows until the page crashes)
function generateWordCloud(canvasId) {
    const ctx = document.getElementById(canvasId).getContext('2d');

    // This object will hold the number of occurances of a technology in the bootcamp
    let techCount = {};

    // Objects are not inherintly itterable in javascript so we must convert the values to an array.
    for (const doc of Object.values(bootcampMetadata)) {

        if ('technolgies' in doc) {
            // we have some top level technologies not associated to an exercise

        }
        // loop over exercises
        if ('exercises' in doc) {
            for (const exercise of Object.values(doc.exercises)) {
                if ('technologies' in exercise) {
                    for (const tech of exercise.technologies) {
                        if (tech in techCount) {
                            techCount[tech] += 1;
                        } else {
                            techCount[tech] = 1;
                        }
                    }
                }
            }
        }
    }

    const data = {
        labels: Object.keys(techCount),
        datasets: [
            {
                data: Object.values(techCount)
            }
        ]
    }

    const options = {
        plugins: {
            wordcloud: {
                fontStyle: 'normal',
                fontColor: '#000',
                fontFamily: 'Helvetica, Arial, sans-serif',
                minFontSize: 12,
                maxFontSize: 50
            }
        }
    };

    const myChart = new Chart(ctx, {
        type: 'wordCloud',
        data: data,
        options: options
    });
}


// Register chart generate on root document only after HTML has been appended to the DOM
// Docsify custom plugin boilerplate found https://docsify.js.org/#/write-a-plugin?id=template
(function () {
    var generateChartOnRoot = function (hook, vm) {
        // Invoked one time after rendering the initial page
        hook.ready(function () {
            //generateWordCloud('wordcloud-canvas');
            generateCategoryDoughnutChart('category-pie-canvas');
        });
    };

    // Add plugin to docsify's plugin array
    $docsify = $docsify || {};
    $docsify.plugins = [].concat(generateChartOnRoot, $docsify.plugins || []);
})();