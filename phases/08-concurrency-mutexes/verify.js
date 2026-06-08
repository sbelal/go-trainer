const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const testFileSource = path.join(__dirname, 'concurrency_test.go');
const testFileDest = path.join(__dirname, '..', '..', 'app', 'internal', 'todo', 'concurrency_test.go');

try {
  // Check if store.go exists in internal/todo
  const storeGo = path.join(__dirname, '..', '..', 'app', 'internal', 'todo', 'store.go');
  if (!fs.existsSync(storeGo)) {
    console.error('  ❌ Error: app/internal/todo/store.go does not exist! Did you finish Phase 7?');
    process.exit(1);
  }

  // Copy test file
  fs.copyFileSync(testFileSource, testFileDest);

  // Run tests with the race detector!
  console.log('  Running go test with -race flag...');
  execSync('go test -race -v ./internal/todo/...', { stdio: 'inherit', cwd: path.join(__dirname, '..', '..', 'app') });
} catch (err) {
  process.exit(1);
} finally {
  // Cleanup test file
  if (fs.existsSync(testFileDest)) {
    fs.unlinkSync(testFileDest);
  }
}
