const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const testFileSource = path.join(__dirname, 'sqlite_test.go');
const testFileDest = path.join(__dirname, '..', '..', 'app', 'sqlite_test.go');

try {
  // Check if sqlite_store.go exists
  const sqliteStoreGo = path.join(__dirname, '..', '..', 'app', 'sqlite_store.go');
  if (!fs.existsSync(sqliteStoreGo)) {
    console.error('  ❌ Error: app/sqlite_store.go does not exist!');
    process.exit(1);
  }

  // Copy test file
  fs.copyFileSync(testFileSource, testFileDest);

  // Run tests
  execSync('go test -v ./app/...', { stdio: 'inherit', cwd: path.join(__dirname, '..', '..') });
} catch (err) {
  process.exit(1);
} finally {
  // Cleanup test file
  if (fs.existsSync(testFileDest)) {
    fs.unlinkSync(testFileDest);
  }
}
