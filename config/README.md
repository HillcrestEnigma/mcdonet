# `github.com/HillcrestEnigma/mcbuild/registry`

## Obtaining the `generated` directory

You must get the latest version of vanilla Minecraft `server.jar` and run the following:

```bash
java -DbundlerMainClass=net.minecraft.data.Main -jar server.jar --server --reports
```

After the command finishes, move the `generated` folder under the same folder as this `README.md` file.
Remove the `generated/.cache` folder.