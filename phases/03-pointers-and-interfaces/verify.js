const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const testFileSource = path.join(__dirname, 'store_test.go');
const testFileDest = path.join(__dirname, '..', '..', 'app', 'store_test.go');

try {
  // Check if store.go exists
  const storeGo = path.join(__dirname, '..', '..', 'app', 'store.go');
  if (!fs.existsSync(storeGo)) {
    console.error('  ❌ Error: app/store.go does not exist!');
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
