async function fetchState() {
  const res = await fetch('/api/state');
  if (!res.ok) return;
  const data = await res.json();
  document.getElementById('count').textContent = data.count;
  document.getElementById('podName').textContent = `pod: ${data.podName}`;
}

async function increment() {
  const res = await fetch('/api/increment', { method: 'POST' });
  if (!res.ok) return;
  const data = await res.json();
  document.getElementById('count').textContent = data.count;
}

window.addEventListener('DOMContentLoaded', () => {
  document.getElementById('incrementBtn').addEventListener('click', increment);
  fetchState();
});
