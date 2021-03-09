package cmd

import (
	"testing"
)

func TestGenerateCmd(t *testing.T) {

	/*
		var stdout bytes.Buffer

		_, err := os.Stat("./julia")
		assert.True(t, os.IsNotExist(err))

		rootCmd.SetArgs([]string{"generate", "1", "--language", "julia"})
		rootCmd.SetOut(&stdout)
		err = rootCmd.Execute()
		assert.NoError(t, err)
		out, err := ioutil.ReadAll(&stdout)
		assert.NoError(t, err)

		assert.FileExists(t, "julia/initenv.jl")
		assert.FileExists(t, "julia/euler001/solution.jl")

		want := ("Language julia saved as default in ./.eulercli_cfg" +
			"Generated the following files for solving problem 1:\n" +
			"   julia/initenv.jl\n" +
			"   julia/euler001/solution.jl\n" +
			"Have fun!\n")
		assert.Equal(t, want, string(out))

		rootCmd.SetArgs([]string{"generate", "--problem", "2"})
		err = rootCmd.Execute()
		assert.NoError(t, err)
		out, err = ioutil.ReadAll(&stdout)
		assert.NoError(t, err)

		assert.FileExists(t, "julia/euler002/solution.jl")

		want = ("Generated the following files for solving problem 2:\n" +
			"   julia/euler001/solution.jl\n" +
			"Have fun!\n")
		assert.Equal(t, want, string(out))

		os.RemoveAll("julia")
	*/
}
