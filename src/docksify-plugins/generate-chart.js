import Chart from 'chart.js/auto';
import { WordCloudController, WordElement } from 'chartjs-chart-wordcloud';
import { bootcampMetadata } from './read-metadata';

// Register the wordCloud controller, element, and scale with Chart.js.
// Starting in Chart.js 3 iirc they support 'tree shaking' which means
// componenets have to be registered explicitly so bundlers like webpack
// can effectively remove unused controllers when minifying
Chart.register(WordCloudController, WordElement);

function generateChart(canvasId) {
    var canvas = document.getElementById(canvasId);
    // var canvas = document.createElement('canvas');
    // container.appendChild(canvas);

    // Set the data for the chart
    var data = {
        labels: ['January', 'February', 'March', 'April', 'May', 'June', 'July'],
        datasets: [
            {
                label: 'My Dataset',
                data: [10, 20, 30, 40, 50, 60, 70],
                backgroundColor: 'rgba(255, 99, 132, 0.2)',
                borderColor: 'rgba(255, 99, 132, 1)',
                borderWidth: 1
            }
        ]
    };

    // Set the options for the chart
    var options = {
        scales: {
            y: {
                ticks: {
                    beginAtZero: true
                }
            }
        }
    };

    // Create a new bar chart
    var myChart = new Chart(canvas, {
        type: 'bar',
        data: data,
        options: options
    });
}

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
                minFontSize: 50,
                maxFontSize: 200
            }
        }
    };

    const myChart = new Chart(ctx, {
        type: 'wordCloud',
        data: data,
        options: {}
    });
}


// Register chart generate on root document only after HTML has been appended to the DOM
// Docsify custom plugin boilerplate found https://docsify.js.org/#/write-a-plugin?id=template
(function () {
    var generateChartOnRoot = function (hook, vm) {
        // Invoked one time after rendering the initial page
        hook.ready(function () {
            // Docsify does not have a typical url structure and the window.location.pathname always appears to be '/' in my testing
            // the page is changed by adding a different hash
            generateWordCloud('chart-canvas');
        });
    };

    // Add plugin to docsify's plugin array
    $docsify = $docsify || {};
    $docsify.plugins = [].concat(generateChartOnRoot, $docsify.plugins || []);
})();