const express = require('express');
const fs = require('fs');
const path = require('path');
const os = require('os');

const app = express();
const PORT = process.env.PORT || 3000;
const DATA_DIR = process.env.DATA_DIR || '/data';
const DATA_FILE = path.join(DATA_DIR, 'counter.json');
const POD_NAME = process.env.POD_NAME || os.hostname();

app.use(express.json());
app.use(express.static(path.join(__dirname, 'src')));
app.get('/', (_req, res) => {
  res.sendFile(path.join(__dirname, 'src', 'index.html'));
});

function readCounter() {
  try {
    if (!fs.existsSync(DATA_DIR)) {
      fs.mkdirSync(DATA_DIR, { recursive: true });
    }
    if (!fs.existsSync(DATA_FILE)) {
      fs.writeFileSync(DATA_FILE, JSON.stringify({ count: 0 }), 'utf8');
      return 0;
    }
    const raw = fs.readFileSync(DATA_FILE, 'utf8');
    const obj = JSON.parse(raw);
    return typeof obj.count === 'number' ? obj.count : 0;
  } catch (e) {
    return 0;
  }
}

function writeCounter(value) {
  try {
    fs.writeFileSync(DATA_FILE, JSON.stringify({ count: value }), 'utf8');
    return true;
  } catch (e) {
    return false;
  }
}

app.get('/healthz', (_req, res) => {
  res.status(200).send('ok');
});

app.get('/api/state', (_req, res) => {
  const count = readCounter();
  res.json({ count, podName: POD_NAME });
});

app.post('/api/increment', (_req, res) => {
  let count = readCounter();
  count += 1;
  if (!writeCounter(count)) {
    return res.status(500).json({ error: 'failed_to_persist' });
  }
  res.json({ count, podName: POD_NAME });
});

app.listen(PORT, () => {
  console.log(`server listening on ${PORT} as ${POD_NAME}`);
});
