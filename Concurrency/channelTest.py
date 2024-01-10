# test_go_with_python.py

import subprocess

def run_go_tests():
    try:
        # Run Go tests
        result = subprocess.run(["go", "test", "-v", "."], capture_output=True, text=True)

        # Print the output
        print(result)

        # Check for test failures
        if result.returncode != 0:
            print("Some tests failed.")
            exit(1)

    except subprocess.CalledProcessError as e:
        print(f"Error: {e}")
        exit(1)

if __name__ == "__main__":
    run_go_tests()
