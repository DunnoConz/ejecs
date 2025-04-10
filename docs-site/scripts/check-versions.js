const fs = require('fs');
const path = require('path');

function checkVersions() {
  const packageJson = JSON.parse(fs.readFileSync('package.json', 'utf8'));
  const packageLock = fs.existsSync('package-lock.json') 
    ? JSON.parse(fs.readFileSync('package-lock.json', 'utf8'))
    : null;

  // Check Node.js version
  const nodeVersion = process.version;
  if (!nodeVersion.startsWith('v18')) {
    console.error('❌ Node.js version must be 18.x');
    process.exit(1);
  }

  // Check critical dependencies
  const criticalDeps = {
    '@cloudflare/kv-asset-handler': '0.3.4',
    '@cloudflare/workers-types': '4.20250410.0',
    'wrangler': '3.114.5',
    'typescript': '5.0.4'
  };

  for (const [dep, version] of Object.entries(criticalDeps)) {
    const installedVersion = packageJson.devDependencies[dep];
    if (installedVersion !== version) {
      console.error(`❌ ${dep} version mismatch. Expected ${version}, got ${installedVersion}`);
      process.exit(1);
    }
  }

  // Ensure no caret or tilde in versions
  const allDeps = { 
    ...packageJson.dependencies, 
    ...packageJson.devDependencies 
  };

  for (const [dep, version] of Object.entries(allDeps)) {
    if (version.startsWith('^') || version.startsWith('~')) {
      console.error(`❌ ${dep} uses ${version}. Please use exact versions without ^ or ~`);
      process.exit(1);
    }
  }

  // Check if package-lock.json exists
  if (!packageLock) {
    console.warn('⚠️ package-lock.json not found. Run npm install first.');
    return;
  }

  console.log('✅ Version check passed!');
}

checkVersions(); 