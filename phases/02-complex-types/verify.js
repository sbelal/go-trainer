const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const testFileSource = path.join(__dirname, 'todo_test.go');
const testFileDest = path.join(__dirname, '..', '..', 'app', 'todo_test.go');

try {
  // Check if todo.go exists
  const todoGo = path.join(__dirname, '..', '..', 'app', 'todo.go');
  if (!fs.existsSync(todoGo)) {
    console.error('  ❌ Error: app/todo.go does not exist!');
    process.exit(1);
  }

  // Copy test file
  fs.copyFileSync(testFileSource, testFileDest);

  // Run tests
  execSync('go test -v ./...', { stdio: 'inherit', cwd: path.join(__dirname, '..', '..', 'app') });
} catch (err) {
  process.exit(1);
} finally {
  // Cleanup test file
  if (fs.existsSync(testFileDest)) {
    fs.unlinkSync(testFileDest);
  }
}
