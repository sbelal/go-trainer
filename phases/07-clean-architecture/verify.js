const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const oldFiles = [
  'basics.go',
  'todo.go',
  'store.go',
  'json.go',
  'sqlite_store.go',
  'server.go',
  'main.go'
];

const newFolders = [
  'cmd/todo-api',
  'internal/todo',
  'internal/sqlite',
  'internal/handler'
];

const testFileSource = path.join(__dirname, 'e2e_test.go');
const testFileDest = path.join(__dirname, '..', '..', 'app', 'cmd', 'todo-api', 'e2e_test.go');

try {
  const appRoot = path.join(__dirname, '..', '..', 'app');

  // 1. Verify old root files are deleted/moved
  for (const file of oldFiles) {
    const filePath = path.join(appRoot, file);
    if (fs.existsSync(filePath)) {
      console.error(`  ❌ Error: Old file "${file}" still exists in app/ root directory. You must remove or move it.`);
      process.exit(1);
    }
  }

  // 2. Verify new folders exist
  for (const folder of newFolders) {
    const folderPath = path.join(appRoot, folder);
    if (!fs.existsSync(folderPath) || !fs.statSync(folderPath).isDirectory()) {
      console.error(`  ❌ Error: Expected clean architecture directory "app/${folder}" does not exist.`);
      process.exit(1);
    }
  }

  // 3. Copy E2E test file to cmd/todo-api
  fs.copyFileSync(testFileSource, testFileDest);

  // 4. Run tests
  execSync('go test -v ./app/...', { stdio: 'inherit', cwd: path.join(__dirname, '..', '..') });
} catch (err) {
  process.exit(1);
} finally {
  // Cleanup
  if (fs.existsSync(testFileDest)) {
    fs.unlinkSync(testFileDest);
  }
}
