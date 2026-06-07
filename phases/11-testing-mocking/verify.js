const fs = require('fs');
const path = require('path');
const { execSync } = require('child_process');

try {
  // Check if router_test.go exists in internal/handler
  const routerTestGo = path.join(__dirname, '..', '..', 'app', 'internal', 'handler', 'router_test.go');
  if (!fs.existsSync(routerTestGo)) {
    console.error('  ❌ Error: app/internal/handler/router_test.go does not exist! You must write your own test.');
    process.exit(1);
  }

  // Run tests in the handler folder
  console.log('  Running user-written unit tests in app/internal/handler...');
  execSync('go test -v ./app/internal/handler/...', { stdio: 'inherit', cwd: path.join(__dirname, '..', '..') });
} catch (err) {
  process.exit(1);
}
