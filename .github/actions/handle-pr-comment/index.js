const core = require('@actions/core');
const github = require('@actions/github');

const github_sha=process.env.GITHUB_SHA
const github_ref=process.env.GITHUB_REF
console.log(`inputs: GITHUB_SHA=${github_sha}`);
console.log(`inputs: GITHUB_REF=${github_ref}`);

try {
    const payload = JSON.stringify(github.context.payload, undefined, 2)
    console.log(`The event payload: ${payload}`);
} catch (error) {
    core.setFailed(error.message)
}
