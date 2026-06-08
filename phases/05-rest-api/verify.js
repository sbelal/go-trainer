const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const testFileSource = path.join(__dirname, 'server_test.go');
const testFileDest = path.join(__dirname, '..', '..', 'app', 'server_test.go');

try {
  // Check if server.go exists
  const serverGo = path.join(__dirname, '..', '..', 'app', 'server.go');
  if (!fs.existsSync(serverGo)) {
    console.error('  ❌ Error: app/server.go does not exist!');
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
