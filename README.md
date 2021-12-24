![Alt text](img/lingue.png)

# lingue
It provides features that standardize some Windows and Linux commands. The cmd cleanup command implementation is available in the package. From systems, Linux and Windows.
</br>
New features will be added as there is a need for implementation in other projects!
</br>

# Example

### 💻 Run cmd cleanup without you worrying about the platform.
____
</br>

```
import (
	"fmt"
	"github.com/Pinablink/lingue"
	"github.com/Pinablink/lingue/oslabel"
)

const dataHeader string = "%s %s %s"

func main() {

	var refLingue *lingue.Lingue = lingue.NewLingue()
	refLingue.ExecCommand(oslabel.CLEAR_CMD)

}

```