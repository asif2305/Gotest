# test_go_with_python.py

import subprocess
import os
def run_go_tests():
    try:
        # Specify the path to the directory containing the "tax" package
        # define_package_path = os.path.abspath("../../Gotest/")
        package_dirs = ["TestWithGo","AnotherGoTest"]
        
        for package_dir in package_dirs:
            package_path = os.path.abspath(os.path.join("../../Gotest", package_dir))
            
                    
            result = subprocess.run(["go", "test", "-v", "./..."], capture_output=True, text=True, cwd=package_path, check=True)
           # result = subprocess.Popen(["go", "test", "-v", "./"], cwd=package_path,stdout=subprocess.PIPE)
            #output, err = result.communicate()
            # Print the output
            #print("re ",result)
            #Process test output
            output = result.stdout
            error_message = result.stderr
            if output:
                print(f"Go test output with package name  -> {package_dir}:\n",output)
                
                passes = output.count('--- PASS')
                fails = output.count('--- FAIL')
                print(f"Passes: {passes}, Fails: {fails}")
            if error_message:
                print("Error message:\n", error_message)
                
            #tests_passed = output.decode("utf-8").count("PASS")
            #tests_failed = output.decode("utf-8").count("FAIL")
           # print(f"Results for package {package_path}:")
            #print(f"Passed: {tests_passed}")
            #print(f"Failed: {tests_failed}")

            # Check for test failures
            if result.returncode != 0:
                print("Some tests failed .")
                exit(1)    
     

    except subprocess.CalledProcessError as e:
        print(f"Error: {e}")
        exit(1)

if __name__ == "__main__":
    run_go_tests()
