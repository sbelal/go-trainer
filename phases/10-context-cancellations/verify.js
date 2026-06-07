const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const testFileSource = path.join(__dirname, 'context_test.go');
const testFileDest = path.join(__dirname, '..', '..', 'app', 'internal', 'todo', 'context_test.go');

try {
  // Check if store.go exists in internal/todo
  const storeGo = path.join(__dirname, '..', '..', 'app', 'internal', 'todo', 'store.go');
  if (!fs.existsSync(storeGo)) {
    console.error('  ❌ Error: app/internal/todo/store.go does not exist!');
    process.exit(1);
  }

  // Copy test file
  fs.copyFileSync(testFileSource, testFileDest);

  // Run tests across the entire codebase to assert signature refactors compile
  execSync('go test -v ./app/...', { stdio: 'inherit', cwd: path.join(__dirname, '..', '..') });
} catch (err) {
  process.exit(1);
} finally {
  // Cleanup test file
  if (fs.existsSync(testFileDest)) {
    fs.unlinkSync(testFileDest);
  }
}
