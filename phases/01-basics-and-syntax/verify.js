const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const testFileSource = path.join(__dirname, 'basics_test.go');
const testFileDest = path.join(__dirname, '..', '..', 'app', 'basics_test.go');

try {
  // Check if basics.go exists
  const basicsGo = path.join(__dirname, '..', '..', 'app', 'basics.go');
  if (!fs.existsSync(basicsGo)) {
    console.error('  ❌ Error: app/basics.go does not exist!');
    process.exit(1);
  }

  // Copy test file
  fs.copyFileSync(testFileSource, testFileDest);

  // Run tests
  execSync('go test -v ./app/...', { stdio: 'inherit', cwd: path.join(__dirname, '..', '..') });
} catch (err) {
  // If exit status is non-zero, it means tests failed
  process.exit(1);
} finally {
  // Cleanup test file
  if (fs.existsSync(testFileDest)) {
    fs.unlinkSync(testFileDest);
  }
}
