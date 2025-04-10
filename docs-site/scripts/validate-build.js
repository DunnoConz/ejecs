const fs = require('fs');
const path = require('path');

// Required files that must exist in the build output
const REQUIRED_FILES = [
  'index.html',
  '200.html',
  '404.html',
  '_nuxt/entry.js',
];

// Required directories
const REQUIRED_DIRS = [
  '_nuxt',
  'docs',
];

function validateBuild() {
  const outputDir = path.join(process.cwd(), 'dist');
  
  // Check if output directory exists
  if (!fs.existsSync(outputDir)) {
    console.error('❌ Build output directory not found!');
    process.exit(1);
  }

  // Check required files
  for (const file of REQUIRED_FILES) {
    const filePath = path.join(outputDir, file);
    if (!fs.existsSync(filePath)) {
      console.error(`❌ Required file missing: ${file}`);
      process.exit(1);
    }
  }

  // Check required directories
  for (const dir of REQUIRED_DIRS) {
    const dirPath = path.join(outputDir, dir);
    if (!fs.existsSync(dirPath)) {
      console.error(`❌ Required directory missing: ${dir}`);
      process.exit(1);
    }
  }

  // Check if index.html is not empty
  const indexPath = path.join(outputDir, 'index.html');
  const indexContent = fs.readFileSync(indexPath, 'utf8');
  if (indexContent.length < 100) {
    console.error('❌ index.html appears to be empty or too small');
    process.exit(1);
  }

  console.log('✅ Build validation passed!');
}

validateBuild(); 