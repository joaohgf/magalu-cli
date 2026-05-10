//go:build unit

package task

import "testing"

func TestUpdateView(t *testing.T) {
	t.Skip("not necessary to test the UpdateView since it embeds the DetailView and all rendering logic is tested in the DetailView tests")
}
