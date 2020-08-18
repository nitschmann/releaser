package branch

// Checkout checks the branch out with the given name using git checkout -b NAME
func (b Obj) Checkout(branchName string) error {
	_, err := b.Git.ExecCommand([]string{"checkout", "-b", branchName})
	return err
}
