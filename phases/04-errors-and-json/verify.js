const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const testFileSource = path.join(__dirname, 'json_test.go');
const testFileDest = path.join(__dirname, '..', '..', 'app', 'json_test.go');

try {
  // Check if json.go exists
  const jsonGo = path.join(__dirname, '..', '..', 'app', 'json.go');
  if (!fs.existsSync(jsonGo)) {
    console.error('  ❌ Error: app/json.go does not exist!');
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
