const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

const testFileSource = path.join(__dirname, 'middleware_test.go');
const testFileDest = path.join(__dirname, '..', '..', 'app', 'internal', 'handler', 'middleware_test.go');

try {
  // Check if middleware.go exists in internal/handler
  const middlewareGo = path.join(__dirname, '..', '..', 'app', 'internal', 'handler', 'middleware.go');
  if (!fs.existsSync(middlewareGo)) {
    console.error('  ❌ Error: app/internal/handler/middleware.go does not exist!');
    process.exit(1);
  }

  // Copy test file
  fs.copyFileSync(testFileSource, testFileDest);

  // Run tests
  execSync('go test -v ./internal/handler/...', { stdio: 'inherit', cwd: path.join(__dirname, '..', '..', 'app') });
} catch (err) {
  process.exit(1);
} finally {
  // Cleanup test file
  if (fs.existsSync(testFileDest)) {
    fs.unlinkSync(testFileDest);
  }
}
